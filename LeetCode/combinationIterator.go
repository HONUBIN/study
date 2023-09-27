package main

type CombinationIterator struct {
	Characters string
	Bits       []bool
	LMB        int
	HasN       bool
}

func constructor(characters string, combinationLength int) CombinationIterator {
	bits := make([]bool, len(characters))
	for i := 0; i < combinationLength; i++ {
		bits[i] = true
	}

	return CombinationIterator{characters, bits, combinationLength - 1, true}
}

func (this *CombinationIterator) next() string {
	word := ""
	for i, bit := range this.Bits {
		if bit {
			word += string(this.Characters[i])
		}
	}
	this.updateBits()

	return word
}

func (this *CombinationIterator) hasNext() bool {
	return this.HasN
}

func (this *CombinationIterator) updateBits() {
	if this.LMB != len(this.Bits)-1 {
		this.Bits[this.LMB] = false
		this.LMB += 1
		this.Bits[this.LMB] = true
	} else {
		if this.LMB == 0 {
			this.HasN = false
			return
		}
		cnt := 0
		for i := this.LMB - 1; i >= 0; i-- {
			if this.Bits[i] && this.Bits[i+1] {
				cnt += 1
				continue
			} else if this.Bits[i] {
				this.Bits[i] = false
				this.Bits[i+1] = true
				for j := cnt; j >= 0; j-- {
					this.Bits[this.LMB-j] = false
					this.Bits[i+2+cnt-j] = true
					if j == 0 {
						this.LMB = i + 2 + cnt
					}
				}
				break
			} else if i == 0 {
				this.HasN = false
			}
		}
	}
}
