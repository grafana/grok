package text

// Defines values for CodeLanguage.
const (
	CodeLanguageGo CodeLanguage = "go"

	CodeLanguageHtml CodeLanguage = "html"

	CodeLanguageJson CodeLanguage = "json"

	CodeLanguageMarkdown CodeLanguage = "markdown"

	CodeLanguagePlaintext CodeLanguage = "plaintext"

	CodeLanguageSql CodeLanguage = "sql"

	CodeLanguageTypescript CodeLanguage = "typescript"

	CodeLanguageXml CodeLanguage = "xml"

	CodeLanguageYaml CodeLanguage = "yaml"
)

// Defines values for TextMode.
const (
	TextModeCode TextMode = "code"

	TextModeHtml TextMode = "html"

	TextModeMarkdown TextMode = "markdown"
)

// CodeLanguage defines model for CodeLanguage.
type CodeLanguage string

// CodeOptions defines model for CodeOptions.
type CodeOptions struct {
	Language        CodeLanguage `json:"language"`
	ShowLineNumbers bool         `json:"showLineNumbers"`
	ShowMiniMap     bool         `json:"showMiniMap"`
}

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	Code    *CodeOptions `json:"code,omitempty"`
	Content string       `json:"content"`
	Mode    TextMode     `json:"mode"`
}

// TextMode defines model for TextMode.
type TextMode string
