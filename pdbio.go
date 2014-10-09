package zpst

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func (p *Protein) WritePDBFile(name string, pdbnumber bool) error {
	f, err := os.Create(name)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var line string
	if pdbnumber {
		for _, r := range p.Residues {
			for _, a := range r.Atoms {
				line = fmt.Sprintf("ATOM  %5d  %-3s%1s%3s %1s%4d%1s   %8.3f%8.3f%8.3f%6.2f%6.2f          %2s%2s\n", a.N, a.Name, a.AltLoc, r.Code, r.Chain, r.Npdb, r.ICode, a.XYZ[0], a.XYZ[1], a.XYZ[2], a.Occ, a.Bfactor, a.Type, a.Charge)
				f.WriteString(line)
			}
		}
		f.WriteString("TER")
	} else {
		for _, r := range p.Residues {
			for _, a := range r.Atoms {
				line = fmt.Sprintf("ATOM  %5d  %-3s%1s%3s %1s%4d%1s   %8.3f%8.3f%8.3f%6.2f%6.2f          %2s%2s\n", a.N, a.Name, a.AltLoc, r.Code, r.Chain, r.N, r.ICode, a.XYZ[0], a.XYZ[1], a.XYZ[2], a.Occ, a.Bfactor, a.Type, a.Charge)
				f.WriteString(line)
			}
		}
		f.WriteString("TER")
	}
	return nil
}

func LoadFromPDBFile(name string) (*Protein, error) {
	var p Protein
	if name == "" {
		return nil, errors.New("You have to enter a valid PDB file name")
	}

	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	data := bufio.NewReader(f)
	var (
		line     []byte
		str_line string

		chain     string
		curNumber int
		resNumber int
		icode     string
		cur_icode string
		nres      int
	)

	p.Name = name
	nres = 1
	for {
		line, _, err = data.ReadLine()
		if err != nil {
			break
		}

		if string(line[:4]) == "ATOM" {
			str_line = string(line)

			fmt.Sscanf(str_line[22:26], "%d", &resNumber)
			fmt.Sscanf(str_line[26:27], "%s", &icode)
			if resNumber != curNumber || chain != str_line[21:22] || icode != cur_icode {

				curNumber = resNumber
				cur_icode = icode

				if chain != str_line[21:22] {
					chain = str_line[21:22]
					p.Chains = append(p.Chains, chain)
				}

				var r Residue
				r.N = nres
				r.Npdb = resNumber
				r.ICode = icode
				r.Chain = chain

				fmt.Sscanf(str_line[17:20], "%s", &r.Code)

				var a Atom
				fmt.Sscanf(str_line[6:11], "%d", &a.N)
				fmt.Sscanf(str_line[12:16], "%s", &a.Name)
				fmt.Sscanf(str_line[76:78], "%s", &a.Type)
				fmt.Sscanf(str_line[30:38], "%f", &a.XYZ[0])
				fmt.Sscanf(str_line[38:46], "%f", &a.XYZ[1])
				fmt.Sscanf(str_line[46:54], "%f", &a.XYZ[2])
				fmt.Sscanf(str_line[54:60], "%f", &a.Occ)
				fmt.Sscanf(str_line[60:66], "%f", &a.Bfactor)
				fmt.Sscanf(str_line[16:17], "%s", &a.AltLoc)
				if len(str_line) > 79 {
					fmt.Sscanf(str_line[78:80], "%f", &a.Charge)
				}

				r.Atoms = append(r.Atoms, a)
				p.Residues = append(p.Residues, r)
				nres += 1
			} else {
				var a Atom
				fmt.Sscanf(str_line[6:11], "%d", &a.N)
				fmt.Sscanf(str_line[12:16], "%s", &a.Name)
				fmt.Sscanf(str_line[76:78], "%s", &a.Type)
				fmt.Sscanf(str_line[30:38], "%f", &a.XYZ[0])
				fmt.Sscanf(str_line[38:46], "%f", &a.XYZ[1])
				fmt.Sscanf(str_line[46:54], "%f", &a.XYZ[2])
				fmt.Sscanf(str_line[54:60], "%f", &a.Occ)
				fmt.Sscanf(str_line[60:66], "%f", &a.Bfactor)
				fmt.Sscanf(str_line[16:17], "%s", &a.AltLoc)
				if len(str_line) > 79 {
					fmt.Sscanf(str_line[78:80], "%f", &a.Charge)
				}

				p.Residues[len(p.Residues)-1].Atoms = append(p.Residues[len(p.Residues)-1].Atoms, a)
			}
		}
	}
	return &p, nil
}
