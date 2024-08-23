package db

type InterfaceOfPostCRUD interface {
	CreateByObject(*Post) error
	FindById(uint) (*Post, error)
	UpdateByObject(*Post) error
	UnsafeDeleteById(uint) error
	SafeDeleteById(uint) error
	FindAll() ([]Post, error)
	FindAllOrdered() ([]Post, error)
	FindByFuzzyName(string) ([]Post, error)
}

type InterfaceOfReplyCRUD interface {
	CreateByObject(*Reply) error
	FindById(uint) (*Reply, error)
	UpdateByObject(*Reply) error
	UnsafeDeleteById(uint) error
	SafeDeleteById(uint) error
	FindAll() ([]Reply, error)
	FindAllByPostId(uint) ([]Reply, error)
	FindAllByUserId(uint) ([]Reply, error)
}

type InterfaceOfUserCRUD interface {
	CreateByObject(*User) error
	GetUserById(uint) (*User, error)
	GetUserByName(string) (*User, error)
	DeleteUserbyName(string) error
	GetAllUser() ([]User, error)
	GetAllUserOrdered() ([]User, error)
	UpdateByObject(u *User) error
	GetUserByFuzzyName(string) ([]User, error)
	GetUserByLocation(string) ([]User, error)
}

type InterfaceOfPetCRUD interface {
	CreateByObject(*Pet) error
	GetPetByName(string) (*Pet, error)
	UpdateByObject(*Pet) error
	GetPetByFuzzyName(string) ([]Pet, error)
	GetAllPet() ([]Pet, error)
	GetAllPetOrdered() ([]Pet, error)
	DeletePetbyName(string) error
	GetPetByKind(string) ([]Pet, error)
}

type InterfaceOfMessageCRUD interface {
	CreateByObject(*Message) error
	GetMessageById(uint) (*Message, error)
	UpdateByObject(*Message) error
	DeleteMessageById(uint) error
	GetAllMessage() ([]Message, error)
	GetMessageByConvId(uint) ([]Message, error)
	GetMessageBySenderName(string) ([]Message, error)
	GetMessageByFuzzyContent(string) ([]Message, error)
}

type InterfaceOfConversationCRUD interface {
	CreateByObject(*Conversations) error
	GetConversationById(uint) (*Conversations, error)
	UpdateByObject(*Conversations) error
	GetConversationByUser1Name(string) ([]Conversations, error)
	GetConversationByUser2Name(string) ([]Conversations, error)
}

type InterfaceOfFileCRUD interface {
	CreateByObject(*File) error
	GetFileById(uint) (*File, error)
	UpdateByObject(*File) error
	GetFileByCreaterName(string) ([]File, error)
	GetFileByFileName(string) (*File, error)
	GetFileByFuzzyName(string) ([]File, error)
	DeleteFileById(uint) error
}
