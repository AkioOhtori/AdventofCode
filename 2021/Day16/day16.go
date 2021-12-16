package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var path = "input.txt" //path to problem input
const PART int = 1
const LITERAL int64 = 4

var bin string
var a string

var Hex2Bin = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111"}

var answer_pt1 int

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func checkPacket(packet string, p int) int {
	// p := 0 //place TODO
	v := packet[p : p+3]

	version, _ := strconv.ParseInt(v, 2, 64)
	fmt.Println()
	fmt.Printf("We're at postion %v and with version %vb = %v\n", p, v, version)
	p += 3

	t := packet[p : p+3]
	packet_type, _ := strconv.ParseInt(t, 2, 64)
	fmt.Printf("Packet type is %vb = %v", t, packet_type)
	p += 3

	answer_pt1 += int(version)

	if packet_type == LITERAL {
		// literal value, read 5 bits at a time
		//need to add bounds checking
		bits := []string{}
		for x := 0; x < len(packet); x++ {
			bits = append(bits, packet[p:p+5])
			p += 5
			if bits[x][:1] == "0" {
				break
			}

			//don't need to parse bits... yet
		}
		fmt.Printf(" which is a literal of length %v.\n", len(bits))
		// fmt.Println(bits)

	} else { //operator
		// length_type := packet[p : p+1]
		// fmt.Println(length_type)
		// p++
		fmt.Printf(" which is an operator.\n")

		if packet[p:p+1] == "0" {
			p++
			l := packet[p : p+15]
			length, _ := strconv.ParseInt(l, 2, 64)
			p += 15
			fmt.Printf("We had unknown number of subpackets of length %v\n", length)
			// for p <= p+int(length) {
			for x := p; p < x+int(length); x = x {
				// fmt.Printf("Going to subpacket %v\n", packet[p:p+int(length)])
				p = checkPacket(packet, p)
			}
		} else {
			p++
			l := packet[p : p+11]
			subpackets, _ := strconv.ParseInt(l, 2, 64)
			p += 11
			fmt.Printf("The length is %vb = %v subpackets\n", l, subpackets)
			// p += 4 - (p % 4) //finish the packet like a good boy (99% this isn't going to fucking work)
			// // check:
			// if (a[(p / 4) : (p/4)+1]) == "0" {
			// 	p += 4
			// 	// goto check
			// }
			for x := 0; x < int(subpackets); x++ {
				fmt.Printf("Going to subpacket %v\n", packet[p:p+11])
				p = checkPacket(packet, p)
			}
		}

	}
	// fmt.Printf("Currently we're at %v/%v; %v %v\n", len(packet[:p]), len(packet[p:]), packet[:p], packet[p:])
	// fmt.Printf("That represents %v %v\n", a[:(p/4)], a[(p/4):])
	return p
}

func main() {
	//Open input file
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	//Load all lines in file into a slice of slices
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// var input [][]int

	//Go through the instructions and convert them to slices of slices, [y],[x]

	for scanner.Scan() {
		//for p := 0; p < len(a)*4; _++
		a = scanner.Text()
		fmt.Println(a, len(a)*4)

		for _, e := range a {
			bin += Hex2Bin[string(e)] //"stock" conversion suppressed zeros so...
		}
		// fmt.Println(bin)
	}
	// for p := 0; p < len(bin); {
	p := checkPacket(bin, 0)
	fmt.Println(answer_pt1, p, bin[p:])
	// break
	// }
}
