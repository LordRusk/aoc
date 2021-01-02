package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// parse
	uData := strings.Split(input, "\n")
	data := make([]int, len(uData)+1)
	for pos, str := range uData {
		data[pos], _ = strconv.Atoi(str)
	}
	sort.Ints(data)

	adaptMap := make(map[int]int)
	for i := 1; i < len(data); i++ {
		adaptMap[data[i]-data[i-1]]++
	}
	adaptMap[3]++ // for builtin adapter

	// fmt.Printf("%v\n%v\n", data, adaptMap) // visualization
	fmt.Printf("%v is the answer.\n", adaptMap[1]*adaptMap[3])
}

const input = `138
3
108
64
92
112
44
53
27
20
23
77
119
62
121
11
2
37
148
34
83
24
10
79
96
98
127
7
115
19
16
78
133
61
82
91
145
39
33
13
97
55
141
1
134
40
71
54
103
101
26
47
90
72
126
124
110
131
58
12
142
105
63
75
50
95
69
25
68
144
86
132
89
128
135
65
125
76
116
32
18
6
38
109
111
30
70
143
104
102
120
31
41
17`
