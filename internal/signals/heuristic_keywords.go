package signals

// heuristicKeywordTable contains the deterministic vocabulary used by the
// prompt-quality heuristics. English entries are matched as tokens, while
// Chinese entries are matched as phrases because Chinese text does not
// normally separate words with spaces.
type heuristicKeywordTable struct {
	CodeActions          []string
	CodeObjects          []string
	ConstraintTerms      []string
	SuccessCriteriaTerms []string
	SpecStructurePhrases []string
	VerificationTerms    []string
	FrustrationPhrases   []string
}

var heuristicKeywords = heuristicKeywordTable{
	CodeActions: []string{
		// English
		"implement", "fix", "debug", "refactor", "update",
		"change", "add", "remove", "create", "write", "test",
		"lint", "compile", "build", "wire",
		// Simplified and Traditional Chinese
		"修复", "修正", "實作", "实现", "實現", "调试", "除錯",
		"重构", "重構", "更新", "修改", "更改", "變更", "添加",
		"新增", "增加", "删除", "刪除", "创建", "創建", "建立",
		"编写", "編寫", "撰写", "撰寫", "测试", "測試", "检查",
		"檢查", "排查", "排除", "解决", "解決", "处理", "處理", "优化",
		"優化", "编译", "編譯", "构建", "構建", "建置", "接入",
		"整合", "串接",
	},
	CodeObjects: []string{
		// English
		"code", "codebase", "repo", "repository", "app", "backend",
		"frontend", "api", "endpoint", "component", "function", "class",
		"module", "package", "schema", "migration", "test", "tests",
		"bug", "error",
		// Simplified and Traditional Chinese
		"代码", "程式碼", "程式", "代码库", "程式庫", "仓库", "倉庫",
		"项目", "專案", "前端", "后端", "後端", "接口", "介面", "端点",
		"端點", "组件", "元件", "页面", "頁面", "函数", "函式", "類別",
		"模块", "模組", "数据库", "資料庫", "迁移", "遷移", "测试", "測試",
		"错误", "錯誤", "问题", "問題",
	},
	ConstraintTerms: []string{
		// English
		"must", "never", "only", "preserve", "keep", "avoid", "require",
		"requires", "constraint", "constraints", "acceptance", "criteria",
		"success", "expected", "output", "format", "verify", "validation",
		"test", "tests",
		// Simplified and Traditional Chinese
		"必须", "必須", "不要", "不可", "仅", "僅", "保留", "避免", "要求",
		"限制", "约束", "約束", "验收", "驗收", "成功", "预期", "預期",
		"输出", "輸出", "格式", "验证", "驗證", "测试", "測試", "检查", "檢查",
	},
	SuccessCriteriaTerms: []string{
		// English
		"success", "acceptance", "expected", "done when", "should result",
		"output", "criteria",
		// Simplified and Traditional Chinese
		"成功", "验收", "驗收", "验收条件", "驗收條件", "完成条件", "完成條件",
		"完成标准", "完成標準", "预期结果", "預期結果", "预期行为", "預期行為",
		"完成后", "完成後",
	},
	SpecStructurePhrases: []string{
		// English
		"acceptance criteria", "success criteria", "requirements", "steps",
		"plan", "scope", "non-scope",
		// Simplified and Traditional Chinese
		"验收标准", "驗收標準", "成功标准", "成功標準", "需求", "步骤", "步驟",
		"范围", "範圍", "非范围", "非範圍",
	},
	VerificationTerms: []string{
		// English
		"test", "tests", "verify", "verification", "validate", "validation",
		"check", "reproduce", "proof", "run",
		// Simplified and Traditional Chinese
		"测试", "測試", "验证", "驗證", "校验", "校驗", "检查", "檢查",
		"复现", "重现", "重現", "证明", "證明", "运行", "执行", "執行",
	},
	FrustrationPhrases: []string{
		// Simplified and Traditional Chinese
		"不工作", "無法工作", "不能工作", "无法修复", "無法修復", "还是报错",
		"還是報錯", "仍然报错", "仍然報錯", "一直报错", "一直報錯", "同样的错误",
		"同樣的錯誤", "坏了", "壞了", "怎么还", "怎麼還", "为什么不能",
		"為什麼不能", "你弄坏", "你弄壞",
	},
}
