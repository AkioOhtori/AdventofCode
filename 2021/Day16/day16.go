package main

import (
	"bufio"
	"fmt"
	"math"
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

func checkPacket(packet string, p int) (pp int, output []int) {
	v := packet[p : p+3]
	version, _ := strconv.ParseInt(v, 2, 64)

	fmt.Printf("We're at postion %v and with version %vb = %v\n", p, v, version)
	p += 3

	t := packet[p : p+3]
	packet_type, _ := strconv.ParseInt(t, 2, 64)
	fmt.Printf("Packet type is %vb = %v", t, packet_type)
	p += 3

	answer_pt1 += int(version)

	switch packet_type {

	case LITERAL:
		// literal value, read 5 bits at a time, 1st bit being status, rest being part of a number
		bits := []string{}
		for x := 0; x < len(packet); x++ {
			bits = append(bits, packet[p:p+5])
			p += 5
			if bits[x][:1] == "0" { //reached the last 5-bit number, stop
				break
			}
		}
		//Compress the nibbles into an actual number
		var o string
		for _, num := range bits {
			o += string(num[1:]) //lop of the status and keep the nibble only
		}
		//Convert and return the final literal
		temp, _ := strconv.ParseInt(o, 2, 64)
		output = []int{int(temp)}
		fmt.Printf(" which is a literal of value %v %v.\n", bits, output)

	default: //otherwise it is an operator and we need to do some pre-math
		fmt.Printf(" which is an operator of type %v.\n", packet_type)

		if packet[p:p+1] == "0" { //15-bit number representing the number of bits in the sub-packets
			p++
			l := packet[p : p+15]
			length, _ := strconv.ParseInt(l, 2, 64)
			p += 15
			fmt.Printf("We had unknown number of subpackets of length %v\n", length)
			//Calculate the subpackets based on length
			for x := p; p < x+int(length); _ = x {
				var o []int
				fmt.Printf("Going to subpacket %v\n", packet[p:p+11])
				p, o = checkPacket(packet, p)
				output = append(output, o...)
			}
		} else { //type == 1; length is a 11-bit number representing the number of sub-packets
			p++
			l := packet[p : p+11]
			subpackets, _ := strconv.ParseInt(l, 2, 64)
			p += 11
			fmt.Printf("The length is %vb = %v subpackets\n", l, subpackets)

			//Calculate the subpackets based on number expected
			for x := 0; x < int(subpackets); x++ {
				var o []int
				fmt.Printf("Going to subpacket %v \n", x)
				p, o = checkPacket(packet, p)
				output = append(output, o...)
			}
		}

		// Now that we have an output, we need to math it based on operator type
		switch packet_type {
		case 0: //SUM
			var o int
			for _, x := range output {
				o += x
			}
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, o)

			output = []int{o}

		case 1: //PRODUCT
			var o int = 1
			for _, x := range output {
				o *= x
			}
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, o)
			output = []int{o}
		case 2: //MINIMUM
			var o int = math.MaxInt64
			for _, x := range output {
				if x < o {
					o = x
				}
			}
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, o)
			output = []int{o}
		case 3: //MAXIMUM
			var o int = 0
			for _, x := range output {
				if x > o {
					o = x
				}
			}
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, o)
			output = []int{o}
		case 5: //greater than
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, "T_T")
			if output[0] > output[1] {
				output = []int{1}
			} else {
				output = []int{0}
			}
		case 6: //less than
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, "T_T")
			if output[0] < output[1] {
				output = []int{1}
			} else {
				output = []int{0}
			}
		case 7: //equal
			fmt.Printf("We did a type %v on %v and got %v\n", packet_type, output, "T_T")
			if output[0] == output[1] {
				output = []int{1}
			} else {
				output = []int{0}
			}
		default:
			fmt.Println("THESUNTHESUNTHESUN") //should never happen
		}
	}
	pp = p //because it can't be both
	fmt.Println()
	return
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

	//load in only input line and convert it from hex to binary
	for scanner.Scan() {
		a = scanner.Text()
		fmt.Println(a, len(a)*4)

		for _, e := range a {
			bin += Hex2Bin[string(e)] //"stock" conversion suppressed zeros so...
		}
	}

	//And away we go...
	p, answer_pt2 := checkPacket(bin, 0)
	fmt.Println(answer_pt1, p, bin[p:], answer_pt2)
	fmt.Printf("The answer to Part 1 is %v and Part 2 is %v", answer_pt1, answer_pt2[0])
}

//EOF
