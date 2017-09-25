package provider

type Zone struct {
	Alias string
	Name  string
}

type Zones []Zone

func (z Zones) Find(alias string) (string, bool) {
	for _, zone := range z {
		if zone.Alias == alias {
			return zone.Name, true
		}
	}
	return "", false
}
