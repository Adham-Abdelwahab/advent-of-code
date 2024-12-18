package main

import (
	f "fmt"
	sc "strconv"
	s "strings"
)

func seven() {
	input := `196487536: 7 3 6 4 393 4 1 93 3 7 9 7
52253144412: 99 5 9 86 816 52 9 2
52182: 5 2 473 1 99 388 2 1
83680255: 703 1 5 2 390 2 4 352 8
43: 4 1 3
9000: 2 6 5 9 8
100587362976: 57 1 4 734 6 1 1 3 8 7 7 6
67930504: 750 6 985 2 40 73 8
874720: 873 9 1 7 20
13771232: 64 230 45 11 466 2
717039285744: 7 3 4 22 257 843 45 38
4443621361: 10 86 409 98 90 13 6 2
3525323623: 376 1 9 6 1 2 8 7 742
7620505: 4 51 253 6 9 5 5 6 2 9 2 5
377033999: 8 227 382 50 84
10332214389: 672 6 7 8 32 5 1 2 6 1 1 8
145843049: 1 4 5 4 7 4 8 8 7 34 46
285061726719: 82 122 6 623 1 847 9
3051902719556: 305 190 1 1 719 5 55
16687440: 4 61 71 56 946
142165716: 8 2 57 3 70 71 7 106 7
26957840: 92 805 91 4
41639685121: 6 7 7 39 3 2 2 7 1 8 704 1
9105412532: 9 6 2 3 750 46 3 255 7 5
184338: 9 8 2 1 6 3 44 2 35 3 1 9
358522: 8 14 4 15 8
22176863: 29 4 1 672 8 66
3590180452: 6 2 31 5 87 19 429
85524672: 10 32 8 61 9 3 4 4 642
1560: 9 6 70 9 4 19 2
9498291487229: 949 8 2 9 148 7 2 19 1 9
28823232: 37 5 6 6 81 720 72 96 2
45662: 87 4 30 2 5
3366345: 28 3 15 86 4
1144: 61 9 64 9 8
489147883: 92 9 1 9 4 41 741 7 4 2 1
62883: 4 9 2 53 503 8
1059480: 914 63 9 91 836 109
4704008218303: 700 672 821 8 2 99 6
4266527: 1 3 2 9 1 3 5 36 8 37 7 7
3193600252: 638 63 9 50 166 83
1261267830: 28 79 221 6 69 86 5
5752870: 94 5 58 1 211 690 107
86389: 5 9 41 2 9 3 2 842 5 500
4878: 41 6 94 402 58
18690: 3 19 89 128 2
5915: 2 96 62 9 35
3700997: 8 33 3 68 9 8 96 6
1463676414034: 38 51 78 38 140 34
135841: 96 6 93 6 6 31 2 566 1 1
177906645: 53 7 1 98 461 5 2 9 6 4 7
2439746: 844 87 2 1 262
27284400: 9 64 5 12 23 195 48
159: 3 10 40 8 97
57276565: 9 615 4 94 2 943 1 565
18106645: 192 41 23 2 8 48
72941486762: 4 3 2 7 6 67 9 4 205 8 1 9
3053: 1 375 1 8 43
35208403: 9 6 3 32 5 99 6 46 712 3
40253347: 4 6 4 1 2 518 4 150 8 2
38661276674: 8 9 179 88 4 15 81 25 7
863121303625: 778 85 121 30 36 24
113702524957: 7 45 836 4 345 2 45 5 8
675916476303: 403 77 3 9 1 85 620 9
394951008: 6 5 42 31 313 8 1 4 4 6
455002: 1 9 6 8 5 6 136 6 4 236 5
106207783: 68 7 854 1 8 64 5 7 2 1
2307675508: 8 32 1 3 8 2 8 3 1 9 9 1
820448: 4 7 201 387 8
1054: 24 205 506 208 6 105
360762674394: 64 3 7 8 53 3 426 3 1 8 8
6955: 27 64 4 43
466507089: 99 5 1 2 4 5 2 942 2 40 7
47132389: 89 5 1 7 4 84 6 24 7 826
23525427: 14 2 9 72 4 5 55 847 2
371202: 3 81 49 5 96 42 13
79924530600: 729 307 558 5 32 2 5 4
9374688: 6 9 8 3 1 3 3 626 1 344 6
36345855: 83 9 87 184 45 853
52361: 9 114 3 17 30 5
120575: 3 130 97 9 1
117097: 1 57 363 703 3 1 42 9 4
808742980: 7 5 1 5 9 9 9 366 9 3 31 5
98162678290: 394 9 2 5 17 9 948 751
217086884: 820 62 427 9 72
8765743364: 5 410 9 6 1 9 717 68 3 2
869677260: 966 308 90 5 8
227813670: 5 30 862 839 9
64900940: 15 57 9 912 97 39 1
11712593: 8 9 1 16 9 19 1 25 86 9
11675969573809: 30 2 4 6 9 883 547 2 11
696655728: 243 43 926 72 1
8400128820930: 8 4 2 50 12 881 6 4 929
7621236148249: 7 511 9 6 8 514 213 49
560964250: 41 63 1 2 1 5 70 55 1 7
64900334: 805 5 1 8 1 26 8 764 7 8
1144: 4 17 714 402 8 1
508392084: 5 8 9 7 82 862
25382656960: 68 3 8 928 4 8 8 5 9 9 9 5
11181: 40 70 85 93 3
994347884004: 9 2 9 61 4 5 3 88 399 6 9
2490: 15 1 29 960 1
566798: 9 7 643 860 58
43736: 2 2 21 9 38 89 6
1526: 49 18 507 719 8 228
9036: 600 64 1 89 12
390605033088: 449 840 3 31 2 6 87 64
4038: 130 4 6 3
166696: 833 5 4 9 5
27454875: 666 4 80 9 5 875
1713037524: 26 259 506 6 778 744
147539698: 37 628 57 417 385 1
93798470: 93 798 4 32 38
516465780: 2 7 5 9 792 31 2 576 1 8
1659152: 16 591 52
55988560892: 2 96 5 245 7 3 116 8 9 2
4651: 7 6 7 2 28
569957874: 885 23 5 4 2 2 1 7 8 7 2 7
27082109: 53 98 600 6 6 90 78 5 6
21069301029: 483 507 4 419 26
1066806827: 762 200 782 193 7
6697057: 2 64 970 56 3
7804992: 25 603 3 519 721 93
5375: 2 18 319 5
52890888: 91 644 8 418 2 7 8 1 8
152570: 86 54 34 583 29
1378102: 2 4 444 146 4 566 7 1 2
1967901210: 24 3 2 4 9 87 3 716 6 45
17012916: 5 8 8 4 2 29 79 69 1 761
2137772: 3 1 95 6 5 1 5 50 20 1 1
331472: 9 438 4 84 8
854364: 8 4 7 7 362 2
1881604: 50 15 735 588 4 4
160851: 643 3 25 26
144: 3 3 9 8
8223653: 465 81 41 14 5 653
151587: 1 9 1 651 19 870 9
131055677: 696 2 1 6 1 2 56 31 6 7 7
454281: 22 68 54 2 573
1208604: 402 8 6 7 3
1077725790: 525 2 4 3 6 95 9
9194691927: 247 94 919 9 44
46: 4 2 1 1 3
20984935: 567 16 37 3 9
136465200: 2 1 96 1 86 2 8 959 8 90
1000: 8 1 1 6 75 6
3064009229: 7 2 796 9 4 5 9 76 2 80 7
4480681: 2 6 3 1 2 28 1 2 1 6 70 11
4074623042: 7 82 70 709 41
9326740896: 1 6 5 69 64 2 604 54 67
450091899: 1 436 8 437 63 9 81 9
31556312: 7 45 55 5 8 1 3
4950: 2 6 8 6 8 4 578 8 4 9 7 1
975572: 7 42 297 67 521
18019: 40 592 701 372 97 1
88319436: 4 4 25 1 8 8 15 8 9 3 2 39
8383083077: 83 771 655 2 77
3705053292700: 8 373 5 885 7 90 2 700
1072: 2 47 5 1 10 82
3186: 4 1 5 122 42 83 10
806520420: 674 312 843 955 4
8080190: 76 36 443 7 93 398
1670297682: 59 982 2 9 26 9 1 16 82
126319573: 125 75 54 22 273 74
13271453: 29 959 3 2 99 5 3 9 8
166984260: 1 669 83 501 756
16344: 7 21 4 7 159 7 4 1 9 9 9 8
94992719246: 48 1 21 943 21 40 94 3
13836040927557: 678 969 234 9 755 7 1
2601360: 2 4 325 9 3 4 4 10 4 2 97
1300828969440: 638 18 253 492 910
22081929: 5 2 1 2 505 7 8 14 5 3 2 9
218037: 2 8 9 589 9 98
1052: 8 7 4 21 807
5562370952: 73 545 9 370 6 7 28 4
1976833: 73 65 5 12 6 24 8 1
556417: 87 5 32 21 9 1
926252: 9 7 12 2 5 4 8 9 4 472 3 4
350875789: 50 401 175 7 88
6000642: 6 593 2 8 593 49
93556: 26 9 4 4 540 17 98 41
378294: 1 797 79 7 6
1068263236: 27 39 6 819 65
870025: 58 3 25 2 26
155880: 2 554 8 432 6
16005132706: 8 21 82 97 2 4 33 1 707
146807552: 146 80 74 84 68
233200: 18 3 3 6 23 880 5
292458792: 3 9 3 874 55 52 910
315534901114: 27 2 8 913 2 6 2 63 8 7 2
18968677252: 1 5 78 1 8 12 2 3 6 27 50
2380188: 6 658 896 4 1 99 4
63241883: 3 629 41 883
1737932850: 103 99 83 2 93 395 83
6164782: 27 908 745 2 9 122
44117780: 1 4 9 9 205 3 3 17 716 5
8746062736: 10 1 159 212 547 5
333833: 2 320 9 168 6 185
241452: 614 89 3 57 6
161985049: 574 5 43 4 556 6
3181196: 55 82 5 54 86 56
2535773: 1 62 3 40 138 175 23
4979629: 4 7 5 7 5 8 2 9 61 42 277
4558393: 822 4 996 274 390
17162495: 9 6 5 8 40 64 39
23451677795: 4 14 1 6 61 620 925 91
75921: 74 947 625 91 258
44037: 48 93 3 3
3064946019: 2 4 8 5 87 1 8 4 8 8 8 132
4534995: 6 630 9 89 79
21330585: 3 97 6 5 4 399 5 9
14872: 80 13 55 14 25 35
352497597: 6 4 56 15 52 70
33743821: 3 55 99 4 6 4 5 16 58
523154: 6 6 11 6 72
2967896923525: 8 9 533 43 3 4 485 6 5 5
27572: 8 5 687 83 9
162802: 395 11 400 16 386
31077906: 407 76 145 90 9
2591250: 67 9 8 94 447 691 6
835298980: 8 3 7 910 9 98 2 6 8 9 1 9
103680: 24 8 45 4 3
415354725: 7 34 1 34 6 71 855
12349238296: 6 6 3 4 8 9 5 7 21 8 296
1338: 118 67 7 30 6 6
82056240: 19 588 824 9 15
14780880226: 2 281 60 108 224
15341205: 2 191 60 7 5
10858393: 6 4 495 462 991 3
984: 96 831 57
7495754096: 535 20 98 724 70 9 4
204214211627: 2 98 8 1 2 9 42 1 11 627
47613: 1 5 90 69 6
651321: 1 649 40 8 913
405164338928: 755 2 8 8 3 99 6 96 8 9 6
3638: 722 4 568 132 44 4 5
699834432268: 2 1 353 5 59 1 1 8 1 3 2 8
206012919155: 743 72 8 9 5 277 7
164295907: 91 416 62 5 7 2 867
1307595: 500 13 5 4 631 158 3
26514010: 44 6 18 158 1 900 1 6 4
3615982563: 382 945 608 256 2
207108: 79 3 222 204 50
6245904569: 6 245 90 372 2 92 755
710: 702 4 7
13551526826: 5 2 924 350 3 101 28
3562505: 395 6 92 61 867
27815475618: 3 4 4 9 9 9 9 91 9 3 6 86
26504: 561 1 7 6 5 6 5 384 4 8
2421648: 553 8 91 3 536
1461059: 7 447 8 322 19 3
1004762: 337 290 93 8 37 2
345010152: 82 4 135 4 470 152
57195340298: 4 1 6 5 150 373 42 820
31562873: 9 57 97 170 507 2
26710: 52 6 83 13 803
82544442: 65 191 213 22 8 443
694251264: 180 2 9 3 6 93 78 8 8 2
1739165: 358 7 694 2
160796392: 8 4 4 4 152 7 26 7 1 9 5
10602882831: 660 1 5 2 9 85 99 4 54 3
471674880: 261 592 120 512 9
518583: 36 15 83 86 197
1017293513: 339 3 2 935 13
36516: 8 9 3 179 4
25763452: 1 350 24 687 71 3 878
17765280: 322 57 4 601 3 6 9 80
605338516: 6 7 57 5 77 9 1 79 7 94 3
6201468825: 9 3 4 8 3 22 134 9 2 3 9 9
4578816: 3 7 28 8 32 88
678736180822: 2 42 727 623 7 80 822
18342806: 66 965 32 9 88
178596: 2 13 4 18 229 85 2 8 1
441431801: 441 431 798
4576830: 29 607 26 4 9
259702324: 9 8 115 255 1 123
65224036: 5 911 4 9 86 79 5 9
232594845644: 332 278 19 7 9 7 1 644
233833021: 3 77 6 76 330 19
85024: 42 512 2
899665359: 24 9 84 9 58 36 70 89
2510008: 5 993 4 8 9 421 5 3 2
15136298: 82 22 88 8 29 7
171259878336: 5 6 62 42 56 942 9 4 8 8
74847090437: 34 22 470 904 36
91160744: 91 153 2 75 35 9
104067072074: 4 7 15 8 8 720 32 3 2 5 7
21314899226: 491 760 816 1 3 7
115041932125: 1 5 1 316 2 47 2 6 8 5 25
82221366516: 93 86 876 651 6
908053: 5 3 6 708 97 3 1
12269457339: 605 8 472 867 20
441173406: 6 4 37 272 2 2 1 9 6 1 4 5
94428666781: 9 442 86 584 3 6 828 8
263664204472: 7 1 5 1 302 74 72 4 472
159599: 6 9 9 129 470
43110144: 2 43 24 8 924
447287553: 447 2 33 9 5 73 4 5 5 6
1918491: 793 8 42 8 9 5 3 3
48632: 5 97 81 45 6
736855997: 600 136 855 9 88 9
765166: 190 47 79 9 4 9 81
1718642: 1 718 63 7 6
4969256: 58 4 8 9 258
48981: 829 59 70
5620: 479 8 71 4 1
26418001: 4 2 375 259 34 1
35139: 5 3 663
560986394272: 31 3 8 5 6 1 993 6 2 9 8 2
26586917: 443 6 6 609 305 2
798756088: 9 166 8 751 605
85184522: 3 936 2 5 4 5 1 3 1 9 33 1
101873: 9 6 5 501 7 367
28102314098: 631 913 2 813 6
5237: 1 86 1 3 6
3196368: 5 8 4 613 64 3 5 3 1 882
48374158734: 715 7 1 670 1 5 8 725 6
818488: 450 93 2 757 9
497545049: 59 22 455 3 5 2 8 9 2 3 6
3198: 70 221 50 57 2 57 1 7
577918451: 20 1 54 61 4 241 226 2
530: 3 25 7 17 9
288287020: 4 8 30 15 80 670 3 4 9
84835: 49 12 139 8 36
2975914404781: 1 307 5 7 9 21 324 7 81
528: 484 5 24 9 6
13413414: 9 52 7 6 8 3 1 76 1 6 59
44985: 390 51 81 627 172 5
21187: 3 6 2 7 828 9 8 1 7 3 40 6
7416301661: 45 1 8 50 867 7 33 6
258423128684: 41 9 1 1 7 589 812 7
135546966: 1 71 4 477 78 966
1494: 3 3 55 7 835
78534456353: 3 6 66 9 14 5 56 6 5 70
141729944432: 7 2 75 4 4 7 3 699 5 4 32
28887951050: 9 3 25 26 907 80 2 251
56800919662: 161 7 8 9 7 5 6 9 6 62
272051160: 27 177 28 1 160
3741976: 9 750 45 8 550 276 50
1819723819533: 8 3 62 7 85 3 8 7 64 4 6 7
457198141: 90 5 94 69 742
5168492: 8 646 492
3187724862: 318 7 72 486 2
14564225: 7 1 5 221 4 41
27156: 82 7 297 70 136
20166351156: 5 188 8 6 6 2 8 4 3 115 4
50266998: 69 36 7 2 572 995
226828: 8 10 1 7 4
103211: 9 4 82 4 28
774092859: 83 58 670 6 4 90 76 2
22684: 7 8 7 1 4
15471647171: 38 9 6 1 631 5 4 9 354 2
54981: 5 8 8 2 9 369
111361746: 720 63 5 6 491
7264746: 66 869 8 89 769
3374705931: 1 7 731 805 6 634 8 9
259399: 2 4 42 4 5 4 6 2 8 5 4 39
24546: 1 7 2 2 2 3 2 2 5 35 4 9
25428452: 7 5 2 4 99 17 711 452
15297954049: 21 1 734 576 47 26
17842917: 2 3 99 75 65 4 1 95 3 3
693838: 599 94 64 10 196
6882152435: 36 75 31 56 7 45 9 3 5
60716309: 1 6 8 55 6 592 739 5 69
106842624: 96 6 2 78 6 76 8 7 4
16058766080: 5 52 5 5 96 2 9 6 2 8 5 4
19292: 4 4 4 1 85 1 9 436 5 7 5 2
3249403: 3 545 21 10 1 970 28
29317383: 5 740 1 7 3 1 158 5 4 6 5
50611: 25 2 60 4 8
1254768: 637 246 50 94 8
3902: 314 60 1 1 75 88
679: 94 97 3 365 9 26 73 5 7
34698038150: 69 396 5 381 53
180903: 718 2 574 9 5
1521: 14 1 21
5924747: 98 65 75 870 6 77
11673: 7 888 32 728 6 7 2 46
85298460548501: 7 7 1 42 7 9 509 8 575 3
19494822950: 57 1 342 822 950
5640: 7 80 94 51 8
536288180364: 9 9 2 625 3 2 5 5 9 5 367
9285: 98 807 20 3 5
28362159: 419 967 7 7
307530305623: 1 1 74 90 67 5 6 1 4 625
3338259555: 74 9 4 5 6 4 1 994 2 5 2 6
80709274806: 7 76 4 1 7 7 315 33 5 1
42360013: 785 59 44 293 908
4210085696: 4 947 851 97 90 1 696
254624814: 7 209 361 174
229103387744: 29 790 338 77 44
11426236: 696 608 9 3 700
215028: 9 993 2 168 99
7428756: 74 287 5 6
41684727: 60 4 916 9 705 3 9 1 3 7
43559795634: 76 555 30 13 8 5 1 69
11110928: 42 3 46 611 126 4
6583910478: 4 94 3 8 2 570 8 8 7 4 4 1
68677558: 7 98 738 7 33 67 268 5
400155: 4 311 54 31 155
664112898: 66 4 112 331 570
24124412: 7 4 44 8 622 93 2 1 5 2 2
66175851562: 8 2 7 1 80 18 4 5 78 2 2
112075: 63 5 485 7 2
135576785135: 889 61 1 5 5 5 82 5 135
191525547: 6 698 26 87 37 3 38 2
467186457574: 63 8 925 986 45 757 4
4582678900795: 6 4 6 5 5 3 3 5 780 5 793
97088223461: 970 88 1 14 1 82 461
587706920968: 1 8 1 8 6 3 3 336 512 8 5
1441077008: 5 75 6 3 7 3 76 9 94 7 8
148083576: 1 5 9 6 6 6 38 2 2 348 9 9
5390: 18 7 6 420 647 2 4 990
9019570: 154 973 30 4 8 2 847
60731336119: 50 604 12 65 36 116
457467745787: 4 5 7 46 77 4 5 787
11658: 377 9 3 7 8
9413604101: 48 5 863 45 1 101
31323641: 44 791 9 4 1
66692130006050: 7 41 18 42 6 625 756 8
195068231: 52 1 93 4 453 8 26 9 7 1
18756976: 1 45 9 4 1 33 51 7 9 9 8 8
624006253: 1 2 9 5 1 1 9 7 8 7 5 350
248847375080: 21 91 25 1 77 120 59 1
9340917020: 701 5 5 5 1 98 9 13 6 19
2688: 12 88 93 798 512
39407: 318 243 9 7 74
278181: 5 3 98 874 83 3
169370: 340 80 40 984 3 6 8
7238695: 8 1 8 1 4 8 1 1 3 6 633 2
171161545313: 9 1 90 8 3 388 6 6 7 12
18174249: 5 9 534 7 2 6 7 6 9 7 1 2
8972: 85 3 3 9 7 2 185 1 2 7 4
384050: 1 84 9 5 17 7 145 1 29 3
1062936: 80 1 3 57 6 37
143658013710: 206 702 177 1 695
57517879168: 7 7 1 357 877 2 837
1032889832: 3 4 86 8 89 832
2752226: 14 565 250 96 8 19
19996: 190 68 1 927 1
96934: 9 2 15 226 4 8 283 6 4
9952: 4 1 6 4 2 264 8 5 125 6 4
245576032: 95 5 100 2 517
261: 3 6 29
5538605: 8 1 2 92 51 3 2 7 340 6
386673: 5 715 571 93 989 98 8
29768773: 7 3 474 329 817 6 1 6
659885: 62 1 38 810 17 58
77998727: 1 23 9 5 3 8 709 9 5 2 2 5
471646091: 589 5 522 54 8 8
229436040: 6 9 3 1 4 4 75 9 6 857 8 5
585474640: 2 982 653 27 298
141867957: 4 860 8 65 633
10482517: 42 5 647 8 22
42395961600: 9 6 20 370 349 4 76
68239589994: 5 23 7 7 37 8 69 491 7 6
3255145: 444 3 706 8 79 978
696306: 272 3 9 4 211 6
74832333323: 74 83 23 33 323
18522726: 209 6 109 84 7 24
70087248: 473 49 84 6 6
4016124: 53 888 741 6 84
145902: 462 56 92 477 496 5 3
45108: 64 3 4 7 4 2 8 61 967 9
60156759464: 287 8 306 1 9 766 209
25622169: 26 6 893 2 6 9 3 2 8 59
899: 1 833 29 30 7
64930028187: 66 1 1 6 2 7 6 535 8 1 86
6813: 673 8 5
31928: 8 62 62 203 9 874 81 9
1406412945: 887 519 412 846 99
204769892: 590 3 907 3 39 257 2
171923: 9 19 85 7 3
658746: 22 60 342 4 5 2
551812: 927 31 9 64 3
488397960: 24 41 9 334 535 29 4 5
11191: 679 14 4 982 698 1
32862667: 6 33 4 19 53 1 72 8 939
21920: 26 843 2
15833965: 89 82 7 7 880 9 2 989
423450951: 163 72 25 18 921 30 1
50149594: 4 497 495 80 17
302470: 316 5 6 16 880 630
157595: 6 670 899 9 5
163371: 2 7 66 291 39 3 2 7 97 6
27361: 842 31 7 386 866
919869: 6 574 4 4 7 741 2 7 7 6
17201061545: 573 368 718 3 5
529162208190: 37 8 3 1 822 5 207 274
1092293: 94 5 59 187 26
39759291: 78 886 504 100 649
987: 471 513 5
619529350443: 5 283 31 8 8 6 35 97 52
90469391818: 90 6 449 245 938 8
66844177093708: 167 11 8 354 5 93 708
523865980: 3 146 3 76 926 407
536946035026: 272 5 5 7 4 43 9 1 197
2719944: 36 677 79 74
99418935: 196 480 506 9 45
3856592: 316 37 475 1 23 7 37
70811648: 677 7 295 64 852 48
918702918151: 1 958 9 8 958 7 7 1 1 50
2597393138: 513 8 9 7 8 5 879
35669086607: 5 889 34 92 235 867 5
49261620097: 6 845 966 8 272 745
76290: 76 283 7
110030: 39 45 6 13 5 10
13887219668497: 789 4 7 537 3 468 90 4
745900: 8 61 60 1 809
5888342598: 25 113 22 4 71 33 6 1 2
85520: 95 9 20
173629: 69 14 963 90 16
13808: 6 4 3 778 30
242556825: 5 688 35 6 824
387127945: 4 41 86 5 7 79 43
53522162: 6 1 792 9 230 6 5 4 12 2
6765390348: 1 9 52 8 71 88 71
560035972: 7 43 35 4 4 397 8 794 4
293565384: 557 5 1 8 6 6 9 729 6 1 4
553648: 2 7 27 85 51 48
48334816: 4 833 48 1 7
436868861: 2 12 312 688 54 7
1201572579: 2 7 4 9 4 79 9 49 1 4 4 72
1695421726: 88 85 98 21 728
47001205: 78 241 94 3 1 1 4 1 5
5101363929649: 42 5 8 7 31 6 3 6 592 4 8
393: 88 65 51 97 92
762957: 730 834 3 6 70 65 80
52949606: 714 389 115 48 86
35688905: 31 7 6 8 2 741 134 8 7 5
7569765: 475 7 5 98 442 7 63
61542120985: 19 42 542 1 20 9 85
871: 4 5 9 6 842 5
633766: 5 5 933 7 6 16 70
20081760: 1 403 9 7 4 3 46 4 2 5
114894157: 92 55 3 78 157
5935839: 5 97 9 922 58
3707607: 46 46 798 99
42924: 85 42 7 4 3
2066376: 50 74 8 46 2 78
162948885582: 72 521 298 58 2 13 2
38898: 82 41 89 569 7 3
363636: 110 9 1 9 364 1
3474498: 82 62 3 683 77
1181170041: 4 78 3 2 598 6 2 2 36 6
37619: 2 9 95 1 845 20 38
281000264: 624 445 45 6 8
16368874: 5 158 688 7 4
88151: 12 11 95 747 5
620768: 2 67 11 8 5 6 6 5 9 3 9 65
14052148: 4 54 547 6 14 6 4 279 8
14620292: 21 118 59 83 9
4000647: 47 4 19 224 5 9
38367063255: 3 64 832 87 859 8 5 5
2285118368: 2 508 7 4 915 57 2 8 4 8
9534: 88 13 47 8 5
8270079: 3 742 111 5 78
581160: 83 5 696
5013993600: 53 16 44 4 340 63 77
8568207045: 85 5 9 7 6 8 4 4 134 4 5
3458322: 6 354 6 2 5 2 9 8 841 9 9
467144167359: 8 3 402 9 3 51 2 7 735 9
1920913371: 96 528 31 199 2
66429: 353 26 280 24 503
113811386281: 8 6 852 4 7 6 3 13 8 8 24
8393: 62 3 3 186 3
184572: 26 620 652 11 80
183430020: 1 6 638 2 9 14 9 3 2 7 9 1
12599009477: 2 9 425 36 79 8 94 8 2 5
68138: 5 7 91 2 5 62
38998697028: 56 36 895 76 1 9 7 873
12177270: 8 145 5 2 8 432 87 54
8730: 3 8 9 30 63 71 4 703 1
145279: 9 3 8 57 75 621 15 8 8 9
155522498: 486 32 24 88 1 8
1497797088482: 498 310 24 385 9 7 4
79977: 6 17 784 9
23976: 3 7 9 4 2 4 8 2 223 9 8
20477635628798: 9 45 43 692 313 798
62833414: 35 7 88 70 7 2
1562652: 375 3 53 78
1233933: 1 70 556 1 45 19 4 8 4
21164: 1 8 3 89 76 2 8 69 562
1609404651: 4 3 3 4 2 351 4 6 1 35 8 5
74748: 9 7 5 6 749
5548020: 904 1 3 3 68 942 3
4887: 158 8 77 1 2 24 1
88982: 6 7 870 1 63 3 7 7
215760160: 73 50 2 70 844
3706: 123 2 3 1 2 9 7 3 1 147
349776320: 5 88 207 6 3 2 4 8 3 8 8 5
33394: 4 926 9 6 52
4841: 5 6 3 16 139 6 3 3 71 6 6
133928945515: 248 60 78 9 192 551 2
280897780929: 539 15 1 203 8 521
54631725: 57 898 1 572 5
85335291: 77 57 754 11
31450: 184 1 5 4 17
294: 225 2 22 6 39
447479: 400 5 237 40 2
90993: 2 73 5 596 997
15747976878: 665 317 812 1 46 79 2
3308852733: 3 308 852 682 54
22666644: 499 7 9 6 3 84
3644107: 24 51 8 7 35 6 4
34844: 60 206 27 55 45 1
132412871531: 6 7 85 9 9 2 4 5 287 9 9 9
92809: 72 1 45 9 85 68 12
69336846888: 1 44 822 71 8 468 8 6 2
1670006546007: 56 1 3 3 9 5 302 5 730 7
7744: 739 35 3
25187: 4 4 1 1 70 6 3 4 1 7 43
525323799000: 572 70 93 297 475
78909600699: 9 971 88 915 700
3651964218: 97 5 74 269 3 49 3 687
36479178: 6 515 7 917 8
674793: 114 7 5 4 6 10 4 623 2 7
29255105: 592 65 56 873 47
2480174: 9 25 7 1 1 1 266 80 705
4324242: 34 26 31 717 8 7
664200230578: 810 41 5 8 4 6 1 5 72 2 2
22436694: 9 587 72 4 94 19
35345387: 277 80 990 5 9 9 1 8 4
6945432229: 1 65 3 8 6 2 8 591 46 5
23800017588: 63 4 3 4 33 849 588
2930952: 4 5 22 9 6 5 99 3 3 2 4 6
1325902454: 5 7 4 8 45 4 55 947 4 55
6696: 93 1 72
1707478747: 185 314 80 921 5
7291222: 217 56 6 22 3
1390427: 7 808 575 4 27
124672577535936: 13 7 91 257 753 5 938
120154018: 329 2 363 9 32 86
26158817598: 2 4 4 86 88 6 634 933 6
75047040: 868 6 6 180 8
6498029748281: 382 3 9 3 238 34 7 280
7578522669: 342 3 41 54 669
278686562: 81 86 188 976 40
509176538: 4 8 9 5 9 1 4 2 176 53 8
24420480: 118 4 602 540 23 840
6740: 8 4 2 8 13 83 1
1665900: 9 2 44 2 7 5 1 3 36 50
131329444: 95 2 5 85 7 24 33 4
9375080: 6 80 35 298 52 5
985635541: 3 5 7 2 2 9 673 2 7 9 60 1
1289692832: 8 9 49 2 2 885 552 1 34
5918: 2 8 12 37 19
84179: 46 91 3 18 2 417
7998: 85 79 806 387 82 8
2285310720: 30 755 41 655 96 44
2530: 8 97 4 6 8
30043: 7 479 60 806 77
5441246: 272 4 248 5 6
2056278432: 44 45 608 59 38 8 1
870282: 740 2 84 6 7
352444: 44 8 8 69 877 653 68
53601574: 5 509 13 5 4 41 6 81 67
95136956082: 5 9 4 6 16 95 5 863 222
1138443520: 8 4 9 2 5 55 2 588 4 64
2432750: 37 41 65 42 7 214
328890: 9 307 64 635 95
193354: 98 3 92 566 8 7 1
14888: 8 22 17 71 835 8
233416399: 9 5 4 15 7 614
33489: 8 4 88 532 9
2605588: 823 7 625 408 4 96
1068755165: 375 285 5 16 4
560611: 934 2 2 6 76
1655429561: 5 9 983 8 3 1 3 1 36 9 7 1
1417632630: 5 7 7 3 1 31 4 66 4 43 3 9
397: 33 286 16 54 8
52665615: 7 7 5 773 519 8 3 4 4 5 9
1253: 55 2 699
34370685203: 6 9 6 10 405 59 1 83
44865: 44 853 9
342574197: 27 41 790 166 5 6
774: 7 65 5 52 6
1809253: 1 5 62 75 7 6 6 96 6 37
2587987392: 467 8 8 6 6 7 40 1 3 8 2 3
26200903: 675 16 673 63 64 3 7
87696: 7 9 86 24 696
97372198: 43 83 63 3 169 3 406
314835308828: 6 8 8 9 2 9 1 8 827 2 4 28
14111847340: 43 98 1 18 44 3 333 8
428858: 5 9 5 67 54 1
2728: 72 45 7 22
85946452: 7 9 6 85 1 1 951 2 3 6 16
22063440921: 210 595 8 571 6 1 921
494893: 440 87 66 567 14 4 3
21795780: 69 3 6 5 1 1 5 6 5 5 5 672
2319634614: 1 319 484 24 9 726
88407660: 1 9 29 31 66 149
2718806: 65 38 8 39 21 7 8 78 38
26188510: 70 256 5 8 1 2 3 33 3 1 8
173110836: 93 79 90 8 202 831 5
39997498449: 8 385 1 935 263 477
150759: 35 1 6 716 239 160
7134700: 2 6 8 402 522 4
12606502: 26 7 45 9 7 19 9 895
4191453: 7 8 356 7 168 66 182 7
1339097405: 6 82 8 6 7 584 6 9 7 653
312622: 3 47 3 408 56 8 7 53
57608: 1 3 1 96 8 7 1 7 7 2 1 205
45018033: 1 89 926 528 9 84
1173314: 36 8 94 6 2 7 1 24 1 9 2
368444: 8 34 1 907 8 29 4 8 7 44
3416: 3 96 81 308 7
903267: 3 86 1 32 67
40600: 2 913 3 3 7 7 5 71 4 5 8
104645223643: 3 43 484 6 3 41 20 47
10364: 8 4 8 5 7 2 1 61 9 3 317 7
178886583: 86 55 37 886 5 80 1
23376672: 340 106 71 314 3 8 6
43366: 99 4 40 2 99 4
69522: 2 31 7 4 3
13887990: 1 2 1 637 4 5 3 773 9 2
2639589: 8 7 62 857 40 29
474485358: 4 8 225 4 2 6 9 37 7 46 6
6201826: 3 9 5 4 2 4 3 8 40 4 232 2
184028203001: 4 3 869 8 2 2 820 299 9
56839409: 5 979 8 807 11 71 4
51620: 84 8 2 50 37 42
37648: 37 60 5 7 36
4164208: 4 180 1 1 75 30 7 163 6
8818159350: 5 9 53 97 8 3 8 3 605 7 3
3476638182: 81 38 81 423 536 645
57148: 8 324 5 22 4 5 3
94127881: 166 1 183 308 3
493755: 9 875 5 6
14989: 9 178 8 2 6
765: 2 3 2 8 8
18963277: 21 3 89 6 277
1589187600671: 91 9 1 616 5 18 5 7 671
90725771: 5 3 3 459 3 6 5 1 2 4 1 71
64058474: 3 9 5 68 5 109 5 1 8 9
12989714686: 2 4 5 66 34 81 32 6
1863: 87 3 5 7 1
56847850: 6 5 6 9 454 72 831 8 4 9
6074694: 8 858 885 52 5
575567: 3 72 444 6 2 45 96
3656323: 978 7 2 464 4
233093619: 670 497 48 7 285
36816271690: 90 381 76 670 407
24992065: 657 583 96 7 38
4463858: 919 193 3 96 5 4
30203841774: 1 4 7 84 681 43 3 7 7 7
911754816: 881 5 308 38 8 84
1287378006: 6 3 86 8 9 10 5 40 3 9
1584: 5 5 4 317 20 993
17444815592: 37 315 6 623 2 777 93
161458902: 716 5 50 1 902
692838857: 871 1 6 79 8 857
313418559: 732 8 455 2 47
128446344471: 307 894 3 156 471
3607721326: 4 5 8 769 5 9 9 7 1 314 9
571279: 608 552 6 82 543 2 14
2329962849: 8 809 119 36 8
2311: 3 8 4 9 4 7 5 248 3 6 4 3
1590318606: 17 58 8 5 5 3 23 7 5 5 6 3
46851824: 48 3 97 8 26
189924065409: 95 2 19 105 65 399 9
3238327295: 5 5 4 45 5 2 7 6 7 9 5 371
197628: 116 567 8 93 4 3
154003: 8 549 6 35 53 20
1866: 44 7 699 720
8972154185400: 8 533 8 454 3 74 7 300
2789: 2 7 12 9 68
107670909: 3 3 8 1 6 4 375 81 6 857
562: 2 2 434 20 84
24835688192: 856 12 562 688 74
29652606: 5 8 9 22 80 9 8 9 39 9 6 2
19614632947: 7 781 1 5 999 9 4 553
143092: 411 4 27 85 58 999
16747162636: 3 7 1 1 9 894 79 1 2 1 1 2
51110: 1 9 9 9 70 80
45252: 4 47 5 6 496
22048: 70 5 31 26 8
1607: 658 247 615 1 84
104905945: 7 247 2 59 9 1 7 6 2 9 5
399456: 876 38 2 2 3
1845180: 765 603 4
7562119: 8 61 75 9 15 9 5 7 9 5 1 9
483897: 4 8 4 49 476 4 9 4 2 6 5 3
119165: 47 85 4 90 3
48932179568280: 9 308 5 7 7 135 12 894
1403892147: 45 5 7 6 4 6 619 8 5 63
454448: 452 2 448
94477: 9 378 2 698
228480166: 81 66 49 7 38 7 6 1 66
1081832: 88 34 7 6 36 9 23
3964471: 7 19 788 97 43
532321464576: 9 709 911 1 578 352 4
9387856: 3 3 38 78 54
1678518074: 6 45 5 9 6 4 30 6 693 4 1
1041: 2 64 234 520 159
1584453: 2 8 32 2 7 151 9 2 4 3 4
53665095: 3 4 1 251 33 906
384620062743: 7 4 68 29 467 9 99 57
442300: 30 55 110 9 6 6 7 32 9 1
232657: 60 7 580 196 5
6859: 23 3 93 1 442
1652: 3 102 90 4 8 60
109060887: 3 10 7 8 5 5 7 8 7 1 7 10
268934: 457 7 12 61 80 7 5
3479112: 253 20 81 9 12 91
44630: 36 9 376 53 2 4
873: 2 8 8 7 4 2 515 7 8 1 68
5848578: 95 72 198 831 2
1244: 6 1 743 10 485
9402174: 10 3 910 288 32 342
54014712581: 53 92 9 47 125 79
57392: 2 87 9 16 9 7 309 6 9 8
5715: 1 67 83 8 63
1663244: 6 28 99 20 1 21
847131299: 5 6 7 92 9 9 9 828 4 2 8 2
134707807567: 9 393 849 5 4 3 478 1
22865746429: 5 9 5 96 6 3 79 3 5 514 8
79177: 23 98 397 700 65 9
2795523: 589 321 384 8 3
91773162834: 4 8 408 6 640 77 6 739
6045777803: 37 812 1 3 9 9 12 69 1 9
947: 17 4 757 9 7
2324896900: 4 7 53 468 280 7 8 9 8 1
519060: 76 23 745 3 205
19585227687: 19 585 2 27 68 4
281990: 87 6 66 821 11 85 2
4550694519: 75 8 4 48 92 2 5 3 6 561
26844551: 8 1 94 757 44 3 7 509 5
200494: 7 346 1 564 50 19 769
190512: 23 1 4 2 81
8972430193: 4 799 8 4 3 4 3 74 7 6 6 1
664: 219 1 417 1 12 16
134438486: 9 49 3 9 3 87 4 87 7 486
113896644887: 867 70 977 8 59 22 22
365757344: 8 36 14 743 7 1 467 5 9
9467902: 946 1 78 23 77
33380171565: 4 4 34 472 32 28 9 4 6
29772900: 85 553 348 456
1071356979: 4 5 1 8 3 998 3 2 4 9 144
83538: 8 3 9 18 39
858023: 814 3 2 7 75 167 6
2070429: 1 11 78 23 429
190642: 9 21 68 96 2
144921: 86 58 2 7 21
258225: 65 6 73 6 6
959650: 5 5 219 1 5 93 7
5796: 64 4 9
1250893526: 3 9 6 8 729 30 81 7 575
424065144009: 3 604 34 43 680 7 173
4322902: 828 9 58 74 2`

	lines := s.Split(input, "\n")

	var total int64
	for _, line := range lines {
		equation := s.Split(line, ":")
		value, _ := sc.ParseInt(equation[0], 0, 64)
		raw_operands := s.Split(equation[1], " ")
		initial, _ := sc.ParseInt(raw_operands[1], 0, 64)
		if calibrate(initial, value, raw_operands[2:]) {
			total += value
		}
	}

	f.Printf("7) %v\n", total)
	f.Printf("7) %v\n", total)

}

func calibrate(accumalator int64, value int64, operands []string) bool {
	if len(operands) == 0 {
		return accumalator == value
	}
	number, _ := sc.ParseInt(operands[0], 0, 64)
	concat := sc.Itoa(int(accumalator))
	next, _ := sc.ParseInt(concat+operands[0], 0, 64)

	mul := calibrate(accumalator*number, value, operands[1:])
	add := mul || calibrate(accumalator+number, value, operands[1:])
	cat := add || calibrate(next, value, operands[1:])
	return cat
}
