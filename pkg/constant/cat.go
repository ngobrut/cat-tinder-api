package constant

type CatRace string
type CatSex string

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

	Male   CatSex = "male"
	Female CatSex = "female"
)

var ValidCatSex = map[string]bool{
	string(Male):   true,
	string(Female): true,
}

var ValidCatRace = map[string]bool{
	string(Persian):          true,
	string(MaineCoon):        true,
	string(Siamese):          true,
	string(Ragdoll):          true,
	string(Bengal):           true,
	string(Sphynx):           true,
	string(BritishShorthair): true,
	string(Abyssinian):       true,
	string(ScottishFold):     true,
	string(Birman):           true,
}

var CatRaces = []string{
	string(Persian),
	string(MaineCoon),
	string(Siamese),
	string(Ragdoll),
	string(Bengal),
	string(Sphynx),
	string(BritishShorthair),
	string(Abyssinian),
	string(ScottishFold),
	string(Birman),
}

var CatSexes = []string{
	string(Male),
	string(Female),
}
