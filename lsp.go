package main

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/#initialize

type InitializeParams struct {
	ProcessID             int                `json:"processId,omitempty"`
	RootPath              string             `json:"rootPath,omitempty"`
	RootURI               string             `json:"rootUri,omitempty"`
	InitializationOptions InitializeOptions  `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities,omitempty"`
	Trace                 string             `json:"trace,omitempty"`
}

type InitializeOptions struct {
}

type ClientCapabilities struct {
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
}

type TextDocumentSyncKind int

type ServerCapabilities struct {
	TextDocumentSync                 TextDocumentSyncKind             `json:"textDocumentSync,omitempty"`
	HoverProvider                    bool                             `json:"hoverProvider,omitempty"`
	CompletionProvider               *CompletionOptions               `json:"completionProvider,omitempty"`
	SignatureHelpProvider            *SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
	DefinitionProvider               bool                             `json:"definitionProvider,omitempty"`
	TypeDefinitionProvider           bool                             `json:"typeDefinitionProvider,omitempty"`
	ImplementationProvider           bool                             `json:"implementationProvider,omitempty"`
	ReferencesProvider               bool                             `json:"referencesProvider,omitempty"`
	DocumentHighlightProvider        bool                             `json:"documentHighlightProvider,omitempty"`
	DocumentSymbolProvider           bool                             `json:"documentSymbolProvider,omitempty"`
	WorkspaceSymbolProvider          bool                             `json:"workspaceSymbolProvider,omitempty"`
	CodeActionProvider               *CodeActionOptions               `json:"codeActionProvider,omitempty"`
	CodeLensProvider                 *CodeLensOptions                 `json:"codeLensProvider,omitempty"`
	DocumentFormattingProvider       bool                             `json:"documentFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider  bool                             `json:"documentRangeFormattingProvider,omitempty"`
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
	RenameProvider                   bool                             `json:"renameProvider,omitempty"`
	DocumentLinkProvider             *DocumentLinkOptions             `json:"documentLinkProvider,omitempty"`
	ColorProvider                    bool                             `json:"colorProvider,omitempty"`
	FoldingRangeProvider             bool                             `json:"foldingRangeProvider",omitempty`
	DeclarationProvider              bool                             `json:"declarationProvider,omitempty"`
	ExecuteCommandProvider           *ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
}

type CompletionOptions struct {
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
	TriggerCharacters []string `json:"triggerCharacters"`
}

type SignatureHelpOptions struct{}

type CodeActionOptions struct{}

type CodeLensOptions struct{}

type DocumentOnTypeFormattingOptions struct{}

type DocumentLinkOptions struct{}

type ExecuteCommandOptions struct{}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/#textDocument_didOpen

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/#textDocument_didChange

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/#textDocument_didSave

type DidSaveTextDocumentParams struct {
	Text         *string                `json:"text"`
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/#textDocument_didClose

type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/#textDocument_completion

type CompletionParams struct {
	TextDocumentPositionParams
	CompletionContext CompletionContext `json:"contentChanges"`
}

type CompletionContext struct {
	TriggerKind      int     `json:"triggerKind"`
	TriggerCharacter *string `json:"triggerCharacter"`
}

type CompletionList struct {
	IsIncomplete bool             `json:"isIncomplete"`
	Items        []CompletionItem `json:"items"`
}

type CompletionItem struct {
	Label               string              `json:"label"`
	Kind                CompletionItemKind  `json:"kind,omitempty"`
	Tags                []CompletionItemTag `json:"tags,omitempty"`
	Detail              string              `json:"detail,omitempty"`
	Documentation       string              `json:"documentation,omitempty"` // string | MarkupContent
	Deprecated          bool                `json:"deprecated,omitempty"`
	Preselect           bool                `json:"preselect,omitempty"`
	SortText            string              `json:"sortText,omitempty"`
	FilterText          string              `json:"filterText,omitempty"`
	InsertText          string              `json:"insertText,omitempty"`
	InsertTextFormat    InsertTextFormat    `json:"insertTextFormat,omitempty"`
	TextEdit            *TextEdit           `json:"textEdit,omitempty"`
	AdditionalTextEdits []TextEdit          `json:"additionalTextEdits,omitempty"`
	CommitCharacters    []string            `json:"commitCharacters,omitempty"`
	Command             *Command            `json:"command,omitempty"`
	Data                interface{}         `json:"data,omitempty"`
}

type CompletionItemKind int

const (
	TextCompletion          CompletionItemKind = 1
	MethodCompletion        CompletionItemKind = 2
	FunctionCompletion      CompletionItemKind = 3
	ConstructorCompletion   CompletionItemKind = 4
	FieldCompletion         CompletionItemKind = 5
	VariableCompletion      CompletionItemKind = 6
	ClassCompletion         CompletionItemKind = 7
	InterfaceCompletion     CompletionItemKind = 8
	ModuleCompletion        CompletionItemKind = 9
	PropertyCompletion      CompletionItemKind = 10
	UnitCompletion          CompletionItemKind = 11
	ValueCompletion         CompletionItemKind = 12
	EnumCompletion          CompletionItemKind = 13
	KeywordCompletion       CompletionItemKind = 14
	SnippetCompletion       CompletionItemKind = 15
	ColorCompletion         CompletionItemKind = 16
	FileCompletion          CompletionItemKind = 17
	ReferenceCompletion     CompletionItemKind = 18
	FolderCompletion        CompletionItemKind = 19
	EnumMemberCompletion    CompletionItemKind = 20
	ConstantCompletion      CompletionItemKind = 21
	StructCompletion        CompletionItemKind = 22
	EventCompletion         CompletionItemKind = 23
	OperatorCompletion      CompletionItemKind = 24
	TypeParameterCompletion CompletionItemKind = 25
)

type CompletionItemTag int

type InsertTextFormat int

const (
	PlainTextTextFormat InsertTextFormat = 1
	SnippetTextFormat   InsertTextFormat = 2
)

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}

// =========================================================
// Common Items
// =========================================================

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type TextDocumentContentChangeEvent struct {
	Range       Range  `json:"range"`
	RangeLength int    `json:"rangeLength"`
	Text        string `json:"text"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}