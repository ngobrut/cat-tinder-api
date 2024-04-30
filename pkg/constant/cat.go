package constant

type CatRace string

const (
	Persian          CatRace = "Persian"
	MaineCoon        CatRace = "Maine Coon"
	Siamese          CatRace = "Siamese"
	Ragdoll          CatRace = "Ragdoll"
	Bengal           CatRace = "Bengal"
	Sphynx           CatRace = "Sphynx"
	BritishShorthair CatRace = "BritishShorthair"
	Abyssinian       CatRace = "Abyssinian"
	ScottishFold     CatRace = "Scottish Fold"
	Birman           CatRace = "Birman"
)

var CatRaces = []string{
	string(Persian), string(MaineCoon), string(Siamese), string(Ragdoll), string(Bengal),
	string(Sphynx), string(BritishShorthair), string(Abyssinian), string(ScottishFold), string(Birman),
}

type CatSex string

const (
	Male   CatSex = "male"
	Female CatSex = "female"
)

var CatSexs = []string{
	string(Male), string(Female),
}
