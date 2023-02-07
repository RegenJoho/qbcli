package qbcli

// This interface should be implemented by prompts, questions and the like.
type QuestionHandler interface {
	//This method is called when the question or prompt is to be asked, and may call other QuestionHandlers.
	Ask() error
	//Should return the possible options, but can return an empty array too.
	GetOptions() []string
	//Must return the name of this Handler, which can be shown directly to the user.
	GetName() string
}

//The type for a callback when some user-code should be executed, for example when the user has answered a question.
type QuestionCallback func(answer interface{})
