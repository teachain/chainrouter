package manager

import "github.com/teachain/chainrouter/stub"

type Zone struct {
	chains  map[string]*Chain
	visible bool
}

func (this *Zone) GetChain(path *stub.Path) *Chain {
	return this.getChain(path.GetChain())
}
func (this *Zone) IsEmpty() bool {
	return len(this.chains) == 0
}
func (this *Zone) GetChains() map[string]*Chain {
	return this.chains
}
func (this *Zone) SetChains(chains map[string]*Chain) {
	this.chains = chains
}
func (this *Zone) SetVisible(visible bool) {
	this.visible = visible
}
func (this *Zone) GetVisible() bool {
	return this.visible
}
func (this *Zone) getChain(name string) *Chain {
	return this.chains[name]
}
