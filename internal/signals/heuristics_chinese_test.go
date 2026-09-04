package signals

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeHeuristics_ChinesePromptQuality(t *testing.T) {
	tests := []struct {
		name string
		in   HeuristicInput
		want HeuristicSignals
	}{
		{
			name: "simplified Chinese code task without structure",
			in: HeuristicInput{Messages: []HeuristicMessage{
				{Role: "user", Content: "帮我修复代码里的登录问题"},
			}},
			want: HeuristicSignals{
				ShortPromptCount:            1,
				UnstructuredStart:           true,
				MissingSuccessCriteriaCount: 1,
				MissingVerificationCount:    1,
				NoCodeContextCount:          1,
			},
		},
		{
			name: "simplified Chinese prompt with criteria and verification",
			in: HeuristicInput{Messages: []HeuristicMessage{
				{
					Role: "user",
					Content: "请修复 src/auth/login.ts 中的登录错误。" +
						"必须保持接口兼容。验收标准：错误凭证返回 401。" +
						"请运行测试验证。",
				},
			}},
			want: HeuristicSignals{},
		},
		{
			name: "traditional Chinese prompt with criteria and verification",
			in: HeuristicInput{Messages: []HeuristicMessage{
				{
					Role: "user",
					Content: "請修復 src/auth/login.ts 中的登入錯誤。" +
						"必須保留介面相容。驗收標準：錯誤憑證回傳 401。" +
						"請執行測試驗證。",
				},
			}},
			want: HeuristicSignals{},
		},
		{
			name: "Chinese code task missing verification",
			in: HeuristicInput{Messages: []HeuristicMessage{
				{
					Role:    "user",
					Content: "请重构前端组件，成功标准是保持现有行为。",
				},
			}},
			want: HeuristicSignals{
				ShortPromptCount:         1,
				MissingVerificationCount: 1,
				NoCodeContextCount:       1,
			},
		},
		{
			name: "ordinary Chinese conversation is not a code task",
			in: HeuristicInput{Messages: []HeuristicMessage{
				{
					Role:    "user",
					Content: "请介绍一下如何在规划会议中讨论技术债务，以及如何制定后续改进计划和跟踪方式。",
				},
			}},
			want: HeuristicSignals{},
		},
		{
			name: "Chinese control prompts are ignored",
			in: HeuristicInput{Messages: []HeuristicMessage{
				{Role: "user", Content: "继续"},
				{Role: "user", Content: "好的"},
				{Role: "user", Content: "请继续"},
			}},
			want: HeuristicSignals{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AnalyzeHeuristics(tt.in))
		})
	}
}

func TestAnalyzeHeuristics_ChinesePromptScoreDelta(t *testing.T) {
	heuristics := AnalyzeHeuristics(HeuristicInput{Messages: []HeuristicMessage{
		{Role: "user", Content: "帮我修复代码里的登录问题"},
	}})
	got := ComputeHealthScore(ScoreInput{
		Outcome:           "completed",
		OutcomeConfidence: "high",
		Heuristics:        heuristics,
	})

	require.NotNil(t, got.Score)
	assert.Equal(t, 94, *got.Score)
	assert.Equal(t, map[string]int{
		"constraintless_first_prompt": 1,
		"missing_success_criteria":    1,
		"code_task_without_context":   4,
	}, got.Penalties)
}

func TestAnalyzeHeuristics_ChineseDuplicatePrompts(t *testing.T) {
	content := "请修复登录页面的鉴权错误并保留现有接口行为，运行测试确认结果"
	got := AnalyzeHeuristics(HeuristicInput{Messages: []HeuristicMessage{
		{Role: "user", Content: content},
		{Role: "assistant", Content: "我会检查登录流程。"},
		{Role: "user", Content: content},
	}})

	assert.Equal(t, 1, got.DuplicatePromptCount)
}

func TestAnalyzeHeuristics_ChineseDistinctPrompts(t *testing.T) {
	got := AnalyzeHeuristics(HeuristicInput{Messages: []HeuristicMessage{
		{Role: "user", Content: "请修复登录页面的鉴权错误并保留现有接口行为，运行测试确认结果"},
		{Role: "user", Content: "请说明订单页面的缓存策略，并给出性能分析和部署建议。"},
	}})

	assert.Equal(t, 0, got.DuplicatePromptCount)
}

func TestCountFrustrationMarkers_Chinese(t *testing.T) {
	msgs := []HeuristicMessage{
		{Role: "user", Content: "这个问题还是报错！！！请重新检查"},
		{Role: "assistant", Content: "我会继续检查。"},
		{Role: "user", Content: "请修复并运行测试。"},
	}

	assert.Equal(t, 1, CountFrustrationMarkers(msgs))
}
