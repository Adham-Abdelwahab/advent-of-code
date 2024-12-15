package main

import (
	f "fmt"
	sc "strconv"
	s "strings"
)

func two() {
	input := `45 47 48 51 54 56 54
76 79 81 84 84
30 32 35 36 38 40 44
72 74 77 78 85
54 55 58 57 60
61 64 67 66 68 67
5 7 6 9 12 14 16 16
82 83 84 82 85 88 91 95
38 40 43 46 49 47 53
23 26 28 31 34 35 35 36
74 75 78 80 81 81 79
11 13 16 17 18 18 20 20
20 21 22 23 23 27
63 64 64 65 68 71 73 79
44 47 49 53 55 56 58 59
81 82 84 87 88 92 89
32 35 38 40 44 44
53 55 59 60 62 64 65 69
53 55 56 58 62 67
57 59 62 68 71 73 74 75
22 24 25 30 33 35 38 36
15 17 20 25 28 30 32 32
83 85 88 94 98
74 77 78 80 87 92
57 55 58 61 63 64 65 67
74 71 73 74 75 73
32 31 34 35 35
35 32 34 36 38 40 42 46
35 32 33 34 36 41
69 67 66 67 70 71 74 75
93 91 89 91 94 93
95 93 92 94 96 96
23 22 21 22 23 24 28
27 24 25 26 28 26 29 36
86 84 86 86 88
2 1 2 3 3 6 9 7
35 33 35 35 36 36
15 13 14 14 16 18 19 23
73 70 70 71 78
21 20 23 25 28 32 33 35
27 25 27 28 32 34 31
40 39 40 44 44
67 64 67 71 75
31 28 29 32 33 37 44
9 8 11 14 17 19 25 28
42 41 46 49 51 50
22 21 26 28 31 33 36 36
20 19 26 28 32
32 29 30 37 40 43 48
20 20 23 25 28 31 32 35
69 69 72 75 76 75
51 51 53 54 54
46 46 47 48 49 53
72 72 75 76 77 78 80 85
53 53 54 52 54
79 79 78 79 78
72 72 75 73 75 75
47 47 49 52 54 51 55
19 19 20 18 21 26
84 84 85 85 88 91 93 94
29 29 29 31 33 36 33
15 15 17 18 21 21 21
61 61 61 64 67 71
62 62 63 64 64 70
64 64 67 71 72 75 77 79
60 60 63 67 70 68
76 76 79 81 85 87 88 88
20 20 23 26 28 31 35 39
29 29 33 35 42
25 25 30 32 35
50 50 52 59 60 62 60
60 60 61 63 68 69 69
81 81 83 86 87 92 96
68 68 73 76 81
21 25 26 28 31 32 35 38
38 42 43 44 45 43
35 39 42 44 46 46
19 23 24 26 28 29 33
54 58 60 63 65 67 73
33 37 38 40 42 39 42 44
65 69 70 67 68 65
9 13 10 13 13
78 82 79 81 83 87
34 38 40 38 40 41 48
15 19 19 21 24
36 40 42 42 39
4 8 9 9 9
52 56 57 60 61 61 65
50 54 56 57 57 60 66
26 30 32 36 37
9 13 16 20 22 23 20
55 59 61 65 65
45 49 50 53 57 60 63 67
74 78 82 84 90
41 45 50 53 56
72 76 82 83 85 82
1 5 11 12 12
72 76 79 81 82 89 93
54 58 64 65 70
66 72 74 77 79 81 84 87
47 53 55 57 60 62 59
71 78 79 81 82 84 84
52 57 59 62 64 65 69
59 64 67 69 71 74 76 83
78 83 85 82 83 85 86 89
27 32 34 37 40 37 35
37 42 40 41 44 44
44 51 54 51 53 57
18 24 27 30 32 31 36
57 63 66 68 68 69 71 72
65 70 70 71 73 72
32 37 37 39 42 42
43 50 52 52 55 58 61 65
20 26 28 28 29 30 33 38
50 55 59 62 64 65 66
44 51 52 56 57 58 61 59
12 17 18 21 25 26 28 28
2 9 11 15 19
61 68 72 73 78
17 22 28 30 33
58 64 67 72 74 71
72 78 79 85 88 90 90
69 74 76 83 85 89
24 31 34 35 36 42 49
22 19 16 14 12 10 9 12
74 72 70 67 66 66
42 40 39 36 32
95 93 92 91 85
70 68 67 66 65 67 66 63
82 79 77 74 77 74 77
74 72 70 71 68 68
97 95 94 96 93 91 87
13 11 9 10 8 2
13 10 8 8 6 4
80 78 75 75 73 70 72
64 63 63 60 60
54 51 50 50 46
35 32 29 28 28 27 21
87 85 81 80 78
27 25 22 18 17 14 17
94 92 88 85 82 79 78 78
19 16 15 11 10 7 3
46 44 42 41 37 34 29
54 52 49 44 43 41 40
29 27 26 20 21
97 95 92 85 85
55 54 52 47 46 42
32 31 28 26 23 21 16 9
43 44 42 40 39 38
14 15 14 12 11 10 13
48 50 47 46 46
90 93 90 87 85 82 78
78 80 77 74 72 69 68 62
32 35 37 35 32 29 28 26
15 16 17 15 14 13 16
62 65 63 65 65
75 77 75 78 74
4 7 5 7 6 1
98 99 96 93 93 90
52 55 55 54 57
76 79 79 78 75 75
75 77 74 74 73 70 66
21 23 22 21 21 19 17 12
79 81 80 76 75
69 71 69 65 63 61 59 60
71 73 71 70 69 65 64 64
59 61 59 56 54 50 46
86 89 85 84 83 77
53 55 54 48 47 44 42 39
13 16 15 14 12 5 6
63 64 63 60 54 52 52
55 56 54 47 43
33 36 33 32 27 22
30 30 28 27 25 24 21 19
14 14 12 11 10 12
34 34 32 31 30 27 24 24
18 18 17 16 12
21 21 19 16 10
26 26 27 24 21 20
68 68 67 66 64 66 63 64
77 77 80 78 76 73 73
71 71 73 71 69 66 65 61
25 25 24 23 21 19 20 15
46 46 46 45 44 41 38 37
73 73 72 70 68 68 67 68
54 54 54 53 51 49 49
47 47 47 46 44 43 39
72 72 70 69 67 67 62
93 93 90 89 85 82 81 78
54 54 53 49 47 48
24 24 20 19 19
93 93 89 87 83
30 30 28 24 22 20 14
97 97 90 88 87
76 76 71 68 67 66 65 66
81 81 75 72 72
50 50 45 42 41 37
75 75 69 67 61
49 45 42 41 38 36
56 52 50 49 46 45 47
31 27 24 22 19 18 18
42 38 37 34 33 30 26
77 73 71 70 67 65 63 58
36 32 29 26 24 23 25 22
46 42 39 41 43
85 81 82 80 78 78
49 45 43 45 41
63 59 62 59 54
91 87 85 85 82
68 64 62 59 59 61
9 5 5 4 4
74 70 68 68 67 63
25 21 19 18 16 13 13 7
69 65 64 60 57 54
68 64 60 59 58 57 60
94 90 89 87 85 84 80 80
28 24 20 18 16 12
26 22 20 16 11
95 91 89 83 80
36 32 29 28 23 24
82 78 75 70 70
58 54 49 47 44 40
52 48 47 45 39 32
98 93 92 90 89 86 84
32 25 23 20 18 15 17
26 19 18 15 14 13 10 10
33 28 26 23 20 18 15 11
84 78 75 74 73 70 69 64
86 79 78 76 74 77 74
73 67 65 62 63 64
90 84 81 83 83
94 87 86 84 87 83
40 33 32 30 33 32 29 24
25 19 18 18 15 14
86 80 79 79 78 77 76 77
10 4 3 3 3
55 48 48 47 43
98 92 91 90 87 87 81
21 16 12 11 9 7 4 2
27 22 19 15 14 11 12
22 17 13 12 10 10
97 90 89 86 82 81 77
99 94 91 87 82
72 67 64 63 62 55 53
63 56 54 53 52 51 44 46
46 39 36 33 31 26 26
48 42 37 35 31
27 20 14 13 11 8 1
15 17 18 21 23 25 26 24
65 68 71 74 77 80 81 81
23 26 28 29 31 32 36
84 86 87 90 92 93 98
89 90 91 94 96 95 96
24 25 26 28 25 26 25
34 37 34 37 37
24 26 29 28 32
72 74 72 73 75 80
91 93 93 94 97
21 23 25 25 22
29 32 35 37 37 37
52 54 54 56 57 61
21 23 23 24 30
16 18 19 23 26 28 29 32
81 84 88 89 91 94 95 92
55 57 61 63 63
77 79 80 82 84 87 91 95
18 20 21 24 28 31 34 40
22 23 26 32 35
29 32 33 35 36 41 38
25 28 29 31 33 34 39 39
1 4 9 12 16
31 33 34 41 44 51
90 87 89 90 92 93 95 96
32 30 31 33 31
85 82 83 85 88 90 90
76 73 76 78 80 81 82 86
26 25 27 30 33 35 38 43
66 64 67 69 71 70 73
56 53 52 54 57 60 59
78 75 73 74 74
7 5 7 5 9
57 55 53 55 58 61 64 71
78 76 77 77 79 82
54 52 52 54 51
53 50 53 53 53
88 86 89 91 92 94 94 98
85 83 83 85 86 92
6 3 7 9 12 14
17 16 18 22 24 26 23
20 17 20 23 24 28 30 30
58 56 58 59 63 67
8 5 9 10 17
77 74 75 77 78 80 87 89
82 81 82 84 85 87 94 92
81 78 79 84 84
60 58 60 61 64 71 75
66 63 65 71 76
62 62 65 68 70 72
22 22 24 26 24
16 16 18 21 21
30 30 33 36 38 42
46 46 47 50 51 52 59
3 3 4 2 4 6 9 11
15 15 12 14 16 17 19 17
32 32 35 34 34
92 92 95 92 93 94 98
73 73 70 73 76 83
18 18 18 21 22
39 39 42 42 43 41
88 88 89 92 92 94 96 96
27 27 27 30 32 36
63 63 65 66 66 72
31 31 35 38 39 40 41
10 10 12 16 17 19 17
38 38 42 43 45 45
57 57 61 62 63 65 68 72
11 11 14 18 21 23 24 31
63 63 69 71 74 75
8 8 13 16 15
22 22 23 26 33 35 35
7 7 10 13 16 21 25
20 20 23 24 29 34
79 83 86 88 91 93
38 42 44 45 48 51 53 51
36 40 42 43 45 45
23 27 29 31 35
69 73 74 76 81
14 18 21 22 25 24 27 30
36 40 37 38 40 42 40
9 13 11 14 16 16
73 77 75 78 82
29 33 31 33 39
23 27 30 30 32 34
57 61 64 65 68 68 69 67
16 20 20 23 24 26 26
23 27 30 30 31 32 36
72 76 76 79 86
53 57 59 60 64 65 67 68
63 67 71 74 71
86 90 91 95 96 99 99
62 66 69 73 77
61 65 69 70 75
23 27 30 36 39 40 41 42
30 34 37 38 39 42 47 46
50 54 56 61 61
21 25 27 32 36
22 26 27 33 36 38 45
29 34 37 39 40 42
31 36 38 39 40 43 46 43
34 39 42 43 46 48 51 51
29 35 36 37 40 43 46 50
23 30 33 34 37 40 43 50
90 95 96 93 95
58 64 66 69 67 66
29 34 37 40 37 40 43 43
24 29 32 35 38 39 37 41
11 18 17 19 20 27
49 56 59 59 62
87 93 96 97 97 96
5 11 11 12 15 16 16
45 50 50 51 53 57
65 71 74 74 81
35 42 46 47 48 51 54
76 81 82 83 87 89 87
2 7 8 11 15 17 17
30 36 39 43 46 50
48 53 54 56 60 63 64 70
75 81 87 88 90 91
27 32 34 36 41 39
81 88 95 98 98
32 38 45 46 49 52 53 57
45 50 52 58 59 66
75 72 69 67 68
16 13 10 9 8 6 6
96 95 94 91 88 87 83
44 41 39 37 35 32 27
34 32 33 30 27
64 61 60 58 56 54 55 58
56 55 54 57 54 51 51
41 39 42 41 39 38 34
26 25 24 26 25 19
21 19 17 17 16
89 86 85 84 83 83 80 83
15 12 12 9 9
52 49 47 47 43
50 49 49 46 43 41 38 31
97 95 94 92 91 87 86
39 37 36 33 32 29 25 26
77 76 73 72 70 67 63 63
22 20 16 14 11 7
34 32 31 28 24 23 16
31 29 26 24 17 15 13 11
81 79 77 70 67 65 66
67 64 58 55 55
22 19 17 14 12 11 6 2
36 35 34 33 31 30 23 17
14 15 12 10 7
15 17 14 11 13
3 6 5 4 3 3
68 71 69 67 64 62 58
34 37 35 32 25
29 32 29 28 25 27 24
20 21 24 22 20 19 21
86 87 84 81 83 82 82
16 19 17 14 13 11 13 9
27 28 31 29 28 23
61 62 62 59 56 55 54
28 31 29 29 26 28
78 79 79 76 73 70 70
18 20 19 17 17 13
40 41 40 40 38 36 35 28
87 88 87 83 81 78
41 44 41 40 39 35 33 35
88 90 89 85 85
25 27 23 22 19 18 14
97 98 95 94 90 87 82
90 93 87 85 84 82 79
87 90 89 86 84 78 81
67 70 65 64 61 61
94 97 94 92 89 83 79
57 59 57 52 46
66 66 63 62 59
39 39 36 35 32 29 32
64 64 62 59 57 54 54
27 27 26 24 21 20 16
41 41 40 38 36 35 28
78 78 76 73 70 68 70 69
54 54 52 51 54 53 50 52
55 55 54 51 54 53 53
30 30 29 32 31 27
97 97 96 99 98 93
86 86 86 83 80 78 75 72
93 93 91 89 89 90
97 97 97 95 95
62 62 60 57 57 56 52
75 75 73 73 71 69 62
50 50 46 44 41 40
80 80 79 75 73 71 68 69
38 38 36 33 29 29
98 98 95 91 87
78 78 75 73 69 68 63
84 84 82 79 76 74 67 66
19 19 16 10 8 11
32 32 25 22 20 19 16 16
23 23 22 21 19 13 10 6
80 80 73 72 71 68 62
81 77 75 73 72 71 70 69
65 61 60 58 61
17 13 10 8 8
90 86 83 81 77
69 65 63 62 61 55
46 42 39 36 38 35 34
58 54 53 55 56
80 76 73 76 74 74
63 59 57 54 53 56 54 50
45 41 43 42 35
23 19 19 16 14
55 51 50 49 49 48 47 48
96 92 90 87 84 81 81 81
86 82 80 77 77 74 70
95 91 90 90 84
44 40 38 37 36 32 29
95 91 87 85 82 85
96 92 91 87 85 85
66 62 59 55 51
40 36 32 31 29 27 22
77 73 71 68 62 61 58 55
87 83 81 79 73 72 71 74
52 48 47 41 38 36 36
88 84 81 76 74 71 67
89 85 82 79 73 70 63
87 80 77 75 72 71
18 12 11 9 7 6 9
98 92 89 88 88
78 73 70 69 66 63 60 56
41 35 32 30 29 23
39 33 31 28 30 29
69 64 65 62 64
59 54 51 52 51 49 47 47
69 62 64 61 57
54 48 45 42 41 43 41 34
60 53 50 48 45 44 44 41
52 46 46 45 48
37 30 27 27 24 24
91 84 84 82 80 76
94 87 86 83 83 76
38 32 28 25 23 20 19 17
17 11 7 6 4 5
65 58 57 55 52 48 47 47
60 55 51 48 44
63 58 55 52 50 46 43 36
37 30 28 22 19
96 89 87 81 78 79
24 18 17 16 14 9 6 6
40 33 28 25 23 21 20 16
67 60 57 55 53 47 45 38
48 50 53 56 57 60 62 59
47 48 50 52 55 57 57
34 35 37 38 41 42 45 49
59 60 61 63 64 66 73
42 45 47 50 49 51 53
18 21 23 25 23 24 23
30 33 32 33 36 38 39 39
48 49 52 55 58 56 57 61
54 57 56 57 62
16 18 19 22 22 24
74 77 80 81 82 82 79
17 19 19 22 22
27 29 31 32 32 33 36 40
30 31 31 33 34 37 38 45
77 80 81 82 84 88 89
56 58 60 61 65 63
86 88 92 95 95
58 59 60 61 65 67 71
55 57 58 62 64 66 73
39 41 42 49 50 52 55
8 11 18 21 22 24 26 24
32 34 40 41 42 42
35 36 38 41 48 51 55
63 65 71 73 76 82
89 88 89 91 92 94
34 33 36 37 35
65 64 65 66 69 70 71 71
24 23 24 26 27 31
10 9 12 13 14 17 20 25
47 44 45 46 47 46 49
51 48 50 47 44
40 38 39 36 38 38
75 73 74 71 73 77
47 46 49 50 47 48 50 56
12 10 10 12 15 17
28 26 27 27 30 27
10 7 10 10 11 13 15 15
85 84 84 85 89
35 32 33 35 35 38 44
70 68 71 73 74 76 80 81
75 72 76 77 79 80 83 82
86 84 85 89 92 95 95
12 9 11 13 16 17 21 25
6 4 6 10 16
36 35 37 40 43 50 53 56
68 65 66 68 71 77 74
85 83 86 93 95 96 97 97
25 23 25 27 34 35 38 42
41 38 45 47 48 54
46 46 47 49 51 52 53 54
30 30 33 34 36 37 36
43 43 45 48 50 51 51
87 87 89 92 96
1 1 4 7 8 11 18
91 91 94 96 99 97 99
48 48 51 54 53 52
73 73 76 77 80 82 80 80
66 66 67 68 70 67 71
72 72 70 71 77
7 7 10 13 13 14
73 73 73 76 79 77
51 51 52 55 57 57 59 59
15 15 16 16 20
70 70 73 76 79 82 82 87
19 19 23 26 27 29
40 40 44 46 43
6 6 9 13 15 18 18
16 16 20 23 25 27 30 34
72 72 75 77 79 83 90
43 43 46 53 54 57 58
5 5 6 11 13 10
10 10 13 15 21 21
29 29 30 37 39 43
71 71 78 81 87
45 49 50 51 52 54
21 25 28 29 32 29
38 42 43 46 46
10 14 15 18 19 20 21 25
10 14 17 18 19 22 23 29
8 12 11 14 17 19 21
58 62 61 62 61
32 36 33 35 35
1 5 6 4 8
69 73 74 75 72 75 81
54 58 59 59 61
54 58 58 59 60 62 59
61 65 65 68 71 74 74
84 88 89 89 93
51 55 56 56 58 59 60 65
10 14 17 21 22
13 17 20 21 25 26 25
45 49 53 54 54
80 84 88 90 91 94 98
19 23 24 26 27 28 32 37
42 46 48 54 55
49 53 58 59 62 65 62
60 64 70 73 75 75
53 57 60 67 68 72
27 31 33 35 41 44 51
20 25 27 28 29 30 33 34
37 42 45 46 44
42 49 51 54 55 58 59 59
36 42 44 47 50 52 56
11 18 21 22 29
77 82 79 81 82
9 14 15 12 9
37 44 41 43 43
30 35 33 35 38 39 43
45 52 53 55 57 59 58 64
26 31 31 34 36 39 41
38 43 44 45 47 50 50 49
10 16 18 18 20 20
18 24 25 28 29 29 31 35
52 59 61 64 67 67 68 75
18 25 26 29 32 33 37 38
21 28 32 35 33
24 31 35 36 38 41 41
47 54 55 59 63
60 66 70 71 74 75 80
65 71 72 79 82 85 87
59 65 66 72 70
33 38 39 42 47 50 50
22 29 30 31 33 34 39 43
5 10 11 18 20 21 26
91 89 88 85 84 85
75 73 71 69 69
23 21 19 17 15 13 12 8
52 50 48 45 43 38
34 33 32 31 32 29 26
41 38 40 38 39
72 70 69 72 72
18 15 17 16 13 10 6
86 85 82 83 77
28 25 22 19 18 17 17 14
95 92 92 91 89 91
80 77 76 76 73 73
23 20 18 18 17 15 14 10
60 58 56 56 49
23 21 17 15 14
33 30 26 25 23 20 23
97 96 93 89 89
51 50 48 45 44 40 36
84 82 80 79 75 72 71 64
76 75 68 67 64 62
97 94 88 86 89
89 87 81 78 77 77
97 94 91 86 83 81 80 76
17 15 13 12 7 2
24 27 25 23 20 19 17 16
28 30 27 24 23 21 18 21
8 9 8 7 6 5 2 2
51 54 51 49 46 42
41 43 41 38 36 31
54 56 55 54 53 56 53
68 71 69 70 67 65 67
6 9 8 9 8 6 4 4
51 53 51 49 47 49 48 44
40 43 45 43 42 36
49 50 50 47 45 42 40 38
4 7 6 4 4 2 3
98 99 99 98 98
53 56 54 54 50
81 82 79 79 77 71
88 89 86 82 80
55 58 56 54 51 47 44 45
58 61 58 57 55 52 48 48
28 29 28 24 23 19
56 59 58 54 47
85 88 86 79 76
95 98 97 90 88 86 88
52 53 50 49 47 42 41 41
60 63 60 54 52 50 46
55 57 56 55 49 48 45 39
54 54 51 48 45 43 41 40
60 60 57 54 53 55
60 60 59 58 55 53 51 51
44 44 41 40 36
81 81 78 77 74 71 70 65
46 46 43 46 44
67 67 66 64 65 63 66
74 74 72 69 66 64 67 67
61 61 60 61 59 58 54
48 48 46 48 47 44 39
25 25 25 24 23
88 88 88 86 88
19 19 19 17 14 11 9 9
70 70 67 65 65 64 60
17 17 14 12 12 11 8 2
99 99 96 92 90
56 56 54 52 48 47 50
80 80 77 73 72 69 69
28 28 26 22 20 17 13
70 70 66 64 58
93 93 88 86 83 81 79 77
40 40 34 31 34
47 47 45 42 39 34 34
89 89 82 79 78 74
98 98 97 96 90 87 84 79
47 43 41 38 35 34 32 29
63 59 58 55 53 50 51
80 76 74 71 68 65 65
47 43 42 41 39 38 37 33
84 80 77 75 74 67
87 83 80 78 81 79
18 14 16 13 16
31 27 30 28 27 27
15 11 8 11 9 6 2
22 18 17 20 17 11
52 48 47 44 44 41 40
39 35 35 32 29 27 26 29
50 46 45 45 43 42 40 40
91 87 85 85 81
56 52 50 50 43
25 21 19 18 14 13
78 74 70 69 67 69
21 17 16 15 13 9 9
59 55 54 51 47 43
96 92 90 88 84 83 82 77
49 45 40 37 35
53 49 43 42 41 44
50 46 44 38 35 35
56 52 46 45 43 41 37
46 42 40 37 30 25
63 58 56 54 51 48 47 45
55 50 49 48 47 44 47
26 21 18 15 15
65 58 57 55 53 49
94 89 87 85 84 81 75
65 59 57 55 56 55
12 5 4 5 6
24 17 20 17 17
17 12 9 7 6 8 7 3
44 37 39 36 30
19 12 12 9 7 4 3
99 94 92 89 87 87 90
18 13 12 12 12
65 60 60 58 54
27 20 17 16 14 12 12 5
22 15 11 10 8
67 61 57 55 57
37 32 29 28 25 21 21
94 87 83 82 79 78 75 71
60 54 51 47 44 38
57 52 49 46 40 38
84 79 78 73 70 72
46 39 36 31 30 30
85 80 79 73 71 67
52 46 45 39 38 37 32
76 78 76 74 69 68 66 66
93 93 92 91 88 86 83 84
2 1 3 5 8 6
77 71 69 67 63 61 58 59
17 10 7 6 6
93 94 95 96 98 99 98
66 71 72 74 78 81
10 13 17 18 21 23 23
48 48 45 43 37 30
26 28 28 29 27
93 91 90 92 95 95
75 72 71 68 66 65 62 56
36 40 42 43 46 45 44
29 33 34 35 35 37 38 39
20 25 26 28 30 32 36 34
2 2 4 5 9 12 12
59 63 66 70 70
95 95 98 96 94 92 87
63 56 54 52 55 49
4 8 11 18 20 17
54 56 53 51 48 50 47
99 96 93 91 90 90 88 88
39 39 41 43 45 49
23 23 29 31 34 35
62 58 55 53 46 41
56 59 57 55 55 54 52 48
47 50 53 54 55 57 58
18 16 14 13 10 9
26 29 31 34 36 39
61 59 57 55 54 52 49
34 35 37 39 41 42 44
54 57 58 59 61 64 66 67
24 23 22 19 16 15
63 65 66 67 70 71 72 73
97 96 93 90 87 84 81
84 82 80 79 76 74 72 71
21 23 24 26 28 30 32 33
77 75 74 72 69 67 66 64
65 67 70 72 75 77 79
21 22 24 25 27 28 30
77 79 81 82 83
52 51 48 45 42
21 20 17 15 14 11 10
50 47 46 43 40 37 34
50 52 53 56 59 61 63
74 75 76 77 80 83
85 86 87 89 90 92 94 96
16 14 13 12 9 7 5 4
92 89 86 85 82 81 78
37 36 33 31 30 27 26 25
55 53 51 50 48 46 44 41
4 6 8 11 14 16 17
31 34 36 37 40
21 22 25 28 29
26 28 30 33 35 38 41
49 46 44 41 40 39 38 36
66 67 69 72 75
72 69 66 65 64 61 59 56
88 91 92 95 98
11 14 16 19 20 23
64 62 61 59 58 56 55 53
82 83 85 86 87 89 92 94
91 88 85 84 82
94 93 90 87 84 82 79
11 14 15 16 19
39 36 35 32 30 27 24 21
79 80 81 84 87
49 50 53 56 59 62
41 40 39 38 37
15 14 12 9 6 5 4 1
9 11 13 15 18 19 20
67 68 71 74 77
86 89 92 93 94 97 98
9 11 13 15 17 19
81 83 85 88 91 92
54 52 50 48 47 45 43
25 24 22 20 19
91 89 86 83 80 79 77 75
22 21 19 18 15 14
64 63 61 60 58 55
14 17 19 22 24 26
52 54 57 59 62 65
11 12 14 17 20 21
90 89 88 86 85 84 81
81 82 84 87 89 92 93
55 53 52 49 47 44 43
28 26 24 23 21 19 18 17
83 80 79 78 76 74 71
95 92 89 86 83
69 72 74 76 77
51 54 57 59 62 64 67
13 14 15 16 19 21 24
47 46 44 43 41 40 37 36
75 74 71 68 67 66 64
30 33 34 36 38 40 41 43
21 22 23 26 29
32 35 36 37 38 40 43 45
11 12 13 16 18 19 20 21
20 18 16 15 13 10 9 8
36 39 42 44 46 47
53 51 49 47 44 43 40
82 80 78 75 74 71 70 67
26 25 22 21 20 19 18 17
62 65 67 68 69 70 71 72
85 84 81 79 78 77 75 74
36 38 41 42 45 48
56 54 53 50 49 48
80 79 76 73 72 70 68 67
40 38 37 34 31
97 96 94 92 89 86 83 80
62 59 58 56 54 52
23 26 28 29 30 33
12 15 17 18 21 22 25
73 70 69 66 64 61
34 32 30 29 26 23
99 98 96 94 92
67 70 72 74 75 76
57 60 62 63 66 69 70
76 74 72 69 68 67 66 63
5 7 10 11 13 15
36 38 41 44 46
94 91 88 85 84 83 82 80
13 15 16 17 19 22 25 26
85 87 88 89 90 93
19 17 16 15 13 10 8 5
15 17 20 23 24 26 29 32
93 92 91 88 85 83 80
27 29 30 31 32 33
68 70 71 72 74 77 80
22 20 17 15 13 12 9
23 20 18 15 14 12
56 53 52 49 47
19 16 14 11 10
32 29 28 27 25 24
98 96 93 90 88
86 88 89 91 92
69 66 65 62 60 59
19 20 23 25 28
56 57 60 63 65 68 70 72
50 47 46 44 42 40 39 38
54 53 52 50 49 46 44
5 6 8 9 12 13 16 19
38 41 43 44 47 50
37 36 35 33 31 30 27 26
73 76 78 81 83 84 85
32 29 28 25 24 21 18 16
84 86 87 89 91 94 95
37 40 43 45 48 50 51 54
91 93 94 96 97
20 22 24 27 30
56 59 60 63 64
58 61 62 65 68 69
43 40 37 36 34 31 28 26
33 36 37 38 39 42
49 47 45 43 42 41 39 38
82 84 87 89 90 93
26 23 20 18 17
71 70 69 67 64 63
67 69 72 73 74
99 97 96 93 92
38 41 42 43 44 45 48
12 15 16 19 22 24
72 74 76 79 82 84 85 87
97 96 93 91 90
55 52 50 48 47 46
16 14 12 11 9 7 6 5
99 96 94 91 88 87 85 83
87 86 84 82 81
13 14 15 16 19 21 23
38 41 42 44 47 48
19 18 17 16 13 10
29 30 32 33 35 38 39
35 34 33 31 30
86 87 88 90 93 95
35 32 31 29 27
68 67 64 61 59
80 77 74 72 71
71 70 69 66 65
15 17 18 19 20 21
42 45 47 50 51
33 31 29 28 27 25 24
73 72 71 69 68 67 65 64
43 46 48 51 53 54
39 42 45 48 51 53
53 52 50 49 47
47 50 51 54 56
26 29 30 33 36 37 40 42
39 36 34 33 30 27 24
65 66 68 70 72 75 76 77
51 53 56 57 59 60 62 63
53 52 50 48 47 45 44
21 20 18 16 13 12 10 8
84 85 88 90 92 93
62 64 67 70 71 73
33 34 37 40 41 42
35 32 31 29 26 25
9 10 13 15 16 18 20 23
94 93 90 87 86 85 83
12 9 8 6 5 3 2
15 18 19 20 22 23 25 28
96 94 93 90 88 86 85
89 90 91 93 94 96
77 76 74 73 70
48 51 52 53 54
32 29 27 26 25
36 39 41 42 45 48 51
48 50 52 53 54 56 59
78 77 75 73 72 71
82 81 78 77 74 73 72 71
14 13 11 9 6 5 3 2
69 72 75 77 80 82 83
75 73 70 68 65 63
91 89 87 85 82 81 78 76
98 95 93 92 89
74 72 69 67 64 62 61
67 66 64 63 61 59
73 71 69 68 66
48 50 53 54 55 58 59
38 41 43 46 47 49 50
25 26 27 30 32 33
62 64 66 67 69 71 72
22 19 18 16 15 13
40 38 36 35 34 32
47 48 50 51 53
26 29 31 33 35 38 39
53 52 50 48 46
57 58 59 62 65
40 42 43 44 45 48 50
10 12 14 17 20
47 46 45 44 42 40
70 71 73 75 76 78 81
8 10 12 15 18
50 51 52 54 56 58 59
78 80 83 85 86 88
89 86 85 84 83 81
52 50 48 45 44 41
54 51 48 47 46
26 28 30 32 35 36
56 59 61 63 66 69 70
27 25 24 21 18
79 80 81 84 87 90 93
24 27 28 30 31 34 37 39
9 10 13 15 17
20 21 23 25 27 29 31
42 39 36 33 31 30 28 27
75 72 70 68 67
38 37 34 31 28
43 46 48 49 51 52
38 39 41 43 45
89 88 85 84 83 82 81
42 39 38 35 33 31
14 13 11 9 7 4 2
39 41 42 43 45 46 48
74 75 78 81 84 87 88 91
65 68 70 71 73
60 63 66 68 71 74 76 79`

	lines := s.Split(input, "\n")
	reports := make([][]string, 0)
	for _, line := range lines {
		reports = append(reports, s.Split(line, " "))
	}
	f.Printf("2) %d\n", process_two(reports))
	f.Printf("2) %d\n", process_with_retry(reports))
}

func process_two(reports [][]string) int {
	count := 0
	for _, report := range reports {
		if safe(report) {
			count++
		}
	}
	return count
}

func safe(report []string) bool {
	var dir bool
	first, _ := sc.Atoi(report[0])
	second, _ := sc.Atoi(report[1])
	if first < second {
		dir = true
	} else {
		dir = false
	}

	for i := 1; i < len(report); i++ {
		first, _ = sc.Atoi(report[i-1])
		second, _ = sc.Atoi(report[i])
		dist := first - second
		if dir {
			if dist > -1 || -3 > dist {
				return false
			}
		} else {
			if dist < 1 || 3 < dist {
				return false
			}
		}
	}
	return true
}

func process_with_retry(reports [][]string) int {
	count := 0
	for _, report := range reports {
		if any_safe(report) {
			count++
		}
	}
	return count
}

func any_safe(report []string) bool {
	if safe(report) {
		return true
	}
	for index := range report {
		next := make([]string, len(report)-1)
		copy(next, report[:index])
		copy(next[index:], report[index+1:])
		if safe(next) {
			return true
		}
	}
	return false
}
