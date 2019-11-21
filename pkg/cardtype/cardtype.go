package cardtype

type Card struct {
	Code        string
	Description string
}

func (c Card) String() string {
	return c.Code + ": " + c.Description
}

var CardsByCode map[string]Card

func init() {
	CardsByCode = map[string]Card{}
	for _, c := range Cards {
		CardsByCode[c.Code] = c
	}
}

var Cards = []Card{
	{
		Code:        "001",
		Description: "Visa",
	},
	{
		Code:        "002",
		Description: "Mastercard, Eurocard",
	},
	{
		Code:        "003",
		Description: "American Express",
	},
	{
		Code:        "004",
		Description: "Discover",
	},
	{
		Code:        "005",
		Description: "Diners Club",
	},
	{
		Code:        "006",
		Description: "Carte Blanche",
	},
	{
		Code:        "007",
		Description: "JCB",
	},
	{
		Code:        "014",
		Description: "EnRoute",
	},
	{
		Code:        "021",
		Description: "JAL",
	},
	{
		Code:        "024",
		Description: "Maestro (UK Domestic)",
	},
	{
		Code:        "031",
		Description: "Delta",
	},
	{
		Code:        "033",
		Description: "Visa Electron",
	},
	{
		Code:        "034",
		Description: "Dankort",
	},
	{
		Code:        "036",
		Description: "Cartes Bancaires",
	},
	{
		Code:        "037",
		Description: "Carta Si",
	},
	{
		Code:        "039",
		Description: "Encoded account number",
	},
	{
		Code:        "040",
		Description: "UATP",
	},
	{
		Code:        "042",
		Description: "Maestro (International)",
	},
	{
		Code:        "050",
		Description: "Hipercard",
	},
}
