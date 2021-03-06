package main

var sample = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

var input = []string{
	"{{[([{{[<{{<[[()()]]((<>())[<>[]])>{{{[]<>}[<>]}<[{}{}][[]{}]>}}<<<<<>()>>((<>())<{}()>)><{{[]()}{()",
	"{<<{({{[<<[<(<<>[]><<>>){<{}()>[{}()]}>]{{[{{}{}}([]())][{[]()>[[][]]]}<(<{}()><<>[]>)>}>>{[(",
	"[[<{<{{{([[{[[{}()]<<>{}>]}{{<(){}>{<>}}{{<><>}{<><>}}}][[<<()[]>(<>())>[(())([]())]>]])<{",
	"{<[([[<<<(<<<<(){}>>{([]<>){[]{}}}><<[()()][{}[]]>[(<><>)[<>]]>>){[<[<()<>>[[]()]]<<()[]>(<>{})>>[(([](",
	"{[<(<<{{[[[[(<{}()>[(){}]){[()()]{<>{}}}]]]]{([(<<[]{}><<>[]>><[{}[]]{()[]}>)])([<({<>[]}({}<>))([()]({}[]))>",
	"(<[{<[{([[(<<<{}<>){[]<>}>(([][])([]()))>[({()[]}<{}<>>)<{{}<>}{(){}}>])(([{{}()}<[]()>]<({}<>)(<>",
	"{[<<((<{([({([()]<()()>)({<>{}}<[][]>)}<[{(){}}](<<>>)>){<[(<>{})]({(){}}<<>{}>)><{<{}()><[]()>}>}",
	"[[([{<<(((((({<>[]}<<>()>))((({}{})<[][]>)({<>()})))[{[[()]{<>[]}]}[(<<><>>{[][]})({<>[]}[{}<>])]])[<{<",
	"(<([[{{<((<<[(<>{})]<[<><>](()[])>>>(<<{[][]}([]()))>)))>}}]])[[[({[(((<([{}<>])[([]{})<()()>]>))(<<[(<><>)<[",
	"{(<([<[[[<<{<<{}[]>>}[{(<><>)[{}<>]]{[[][]]<[]<>>}]>[<[<()[]>[{}()]]<([]<>){<>{}}>>[{{()[]}([])}[[[]]]]",
	"({{<{(<(({[{{<<>()><()<>>}}]}){[{{(({}<>>)}}<{<<{}[]>>}>]})((<((((<><>)[(){}]){<(){}><{}()>}){({(",
	"[(<{((<[{({<[<{}><[][]>]<{[]<>][[]()]>>[[[[]{}]](<{}<>><{}>)]}({{{[]()}[[][]]}[<{}[]>[{}[]]]}{(({}",
	"(((<<[{{{{{{<[<>][<>[]]>)}}}<<[[[[[]{}]<[]{}>](<[]{}>)][(<<>{}>({}()))]]([[{{}{}}[{}<>]]{{[]",
	"([<[<({(({<{[{()[]}][({}[]){<>{}}]}>{{{[[]{}][<>()]}}<(({}){()()})<[()<>][()]>>}}<([({<><>}({}))({()()}[",
	"<[{<[{{<{(<{<(()[]){{}{}}><{[]{}}{{}()}>}<{[[][]][<><>]}<<()}({}())>>><([[[][]](()<>)])>)(<([<[]<>><<>[]>]{",
	"<((({{<<([[<{({}<>)({}<>)}{[<><>][(){}]}>({[[]][()<>]}<([]<>){<><>}>)]][{{{{<>[]}[{}()]}[(()())<[][]>]}{",
	"<([<[[(<(((({[{}[]][[]<>]}<([][])<[]{}>>){[(()())({}{})](<<>{}>)})))))[{[({[[{{}{}}({}())]][",
	"([[(([{([[<<{[[]()]{[]{}}}{{<>()}}><[<()()>([]<>)]{<(){}>((){})}>>({(<()<>>{()[]})}[(<[][]><[]<>>)]>]{",
	"(({{[([{({{([<()<>>{{}[]}]{({}<>)(()[])})}((<[[]]<{}[]>>{(()<>)[{}<>]})[(<()()>){{()[]}({}()]}])}{(([",
	"[[[(<[<{{([[[({}[])]({(){}}{<>})]<[<{}>]>](([[<>()]{{}{}}]{<<>()>[<>[]]})((({}())<{}()>)<<[]{}>>))]}(<<(",
	"{{([<[<<({({[[{}[]]<[]()>]<{<>{}}[[]<>]>}(([{}{}]{{}<>})[<[][]><[]{}>]))[[[{()[]}{<><>}]]]}<",
	"{([<[<[{[(<{([()()]<[]<>>}<[()()]<()()>>}<{[[]()]<(){}>}([<>[]]{<>()})>>[[((<>[])(<><>))[{{}()}[<>[]]]]<",
	"<{<{((({[[[<{<[]{}><()[]>}>([<<>[]><<>()>]<[<>()][[]<>]>)]]]}(<({<<[(){}]><{[]<>}[{}[]]>><{{[]{}}((){})",
	"[{{{([{[([<[({[]{}})<({}<>){<>{})>][{{<>}<<>()>}<[<>{}]{()[]}>]>{{[[{}<>][<><>]]{<{}()><[][]>}",
	"<<{<<([[([[(<(()<>){[]()}><{<>[]}<<>[]>>)]{<(<[]{}>((){}))([[][]]({}[]))>{[({}[])(<>)]({()()}<<>[]>)}}]",
	"[[([{(([(<[({({}[])<{}[]>})(<(()())(()())>[{<>{}}({}[])])]>[(<<(<>{})>[<{}{}>[<><>]]><<[()()]",
	"[[{{<<((<[{{<(<>[])(())>}}<{{<()<>>}}>]{{<{<[]>(()[])}(<()()><(){}>)>}(<{<{}[]>(<>())}[<()()>{<>()",
	"<<({<<([({((({{}{}}<[]()>)([{}{}])))[<[<<>{}><{}[]>}{((){}){{}{}}}>(<((){})<{}[]>>([<><>](<>[])))]})[",
	"{({<(([[<{{<[<()>][[[]{}]{<>{}}]>[[[()[]]({}<>)][{[][]}{[][]}]]}<<{[()[]][()()]}[({}{})(<>{})]>)}>{[[[{",
	"(({[[<(<[((<(({}[])([]()))<(<>[])[()<>]>>))<[[[{{}[]}<()[]>}]<{{<>[]}{(){}}}>]>]{[{<<<{}<>",
	"<{{<<[(([[([<{[][]}<<>()>><<<>>>])<{((<>{})[[][]])({{}<>}<[]{}>)}(<<{}{}><{}{}>><(<>[])({}<>)>)>]]{<<",
	"{([[<<[({[<(<(()())><{[]()}<()[]>>)<<([][])(()[])>[<[]<>><<>()>]>><{([<>[]]{{}})<{{}[]}{<><>}>}[{{[]}{[]{}",
	"<{[<<(<[[<{(<{{}[]}<[][]>>{[(){}]})([({}{})<{}()>]<<()><{}>>)}>({<{([][])}<{[]{}}<<>{}>>>{(<[]",
	"{{<[[<([[[{[[{<>[]}<(){}>]([<>[]]({}<>])]{<[{}()]{{}[]}>[<{}()>(()[])]}}<{[((){})]}([({}<>){<><>",
	"([<([{<(<[{[{{{}<>}[()<>]}<<()>{[]<>}>]}[[[<<><>><[]<>>]{<{}[]>{<>{}}}]]](<([<<><>><()()>][({}",
	"{{{[{[{([([<[[{}<>]{{}[]}]{<<><>>{{}<>}}>][<<[[]{}][()()]>{{[]}{{}{}}}>{[<[][]>[{}<>]]{(()<>){[]",
	"[<{{<[[[[<{<({(){}}[<><>]>({{}}<<><>>)><({{}<>}[[][]])(([][])(<>{}))>}>]<[(<{[{}()]{()()}}<((){})>",
	"{<<{{<{<({[{{({}())<{}{}>}[[{}[]](<>[])]}<({[][]}<{}<>>)<<<>[]>[{}[]]>>]{{<{[]<>}<[]<>>>{<<>[]>}}}}<",
	"{[(<({([([{[{[<>[]][()()]}<{{}[]}[{}{}]>]}([{{{}<>}([]())}<[<>()](()[])>]{[(()<>)(<><>)]})]<[([{[]",
	"({((<<({{([[<{(){}}<()[]>>(<[]{}>{<>[]})]<{{{}()}(<>[])}[<()[]><{}()>]>]<<[[{}[]]{{}[]}]({()()}[{",
	"{<{<(((<({[<{{(){}}[<><>]}<<{}{}><{}{}>>>]({<[[]{}](<>)>{[[][]]{(){}}}}[{(())<{}[]>}{<<>[]]{()",
	"<<{<({((<[((<<()<>>[<>[]]>{{{}<>}({}<>)})([(()()){[]()}]))[[[<[]<>><<>>]{<<>()><()[]>}]]]>)[(<{(<{{}{}",
	"[({<{{<<{<<{<[<>[]]<()<>>>[({}())([][])]}>>}>>}((([{{<([()[]]{{}()})<<(){}>(<><>)>><(<()()>(<>{}",
	"{[{<<[(<<[(<<{<>[]}{[]{}}>[[[]{}](()())]>)(([(<>[])(<><>)](<[]<>>(()())))[(({}{})([]{}))])](<<({{}()}",
	"<<((({((({<{{{{}{}}[[]()]}([()<>]{<>()})}>}){<{{<(<>())><<()()>[()()]>}}>{[{([()<>]<[]<>>)[<{}(",
	"(([[[{{[{{<<(<{}>{{}()}){{[]()}{<><>}}><[({}[])<<>{}>](<<>{}>[{}[]])>>([<<[]{}>[()<>]>[([]()){(",
	"<({<{<[({{(([{{}()}[[][]]]<{()<>}{<>()}>))<[[({}())({})][{<><>}<[]()>]]{[[<>]{[]{}}]}>}[({{<{}<>>}}",
	"[({{[<[[[[({<{<>[]}{{}{}}>{{{}{}){()[]}}}(<<()<>>{[]{}}>))]({[{<[]()>(<>[])}[(<><>)]]})]]]{<[<(",
	"<<<[({{{[{[(<{()}(()[])><([][])([][])>)({{{}}{<>{}}})]{<[([]<>)[{}()]]((<>())(<><>))>>}{<{{",
	"({([<(([({<<<<[]()>(<>{})><<[]>{<><>}>>{<{<>}{{}()}>({<><>}{[]{}})}>}[<<{(()<>)[{}<>]}([()()]<()()>)>[",
	"[<({{<{<[{{{(((){}))({{}()})}}({[[{}<>]<<>{}>](((){})[[]<>])})}(<<<<()()><[]<>>><<[]{}><()<>>>><<<()<>><<>()",
	"{{(([<<{{({[<{[]}[<>[]]>(<<><>>[[]{}])]{{(<>())({}())}[<[]()>[()[]]]}}[(<[{}()](<>{})><<{}",
	"{<({{([({<<[[[[]{}]{<>()}]{<(){}>{()()}}]>(<{{(){}}({}{})}<<{}()>>>{[<<>()>>{[<>{}]{<>{}}}})>}",
	"{{([<<{<[<([{[{}[]]}(<<>{}>(<>{}))]<<{()()}[<>[]]>{<[]()>([]{})}>)({(<[]><[]>)[<[]<>><{}()>]}(",
	"<((<(({({[(<{[<>[]][{}{}]}{{(){}}}>{{<<>[]>{[]}})){<{<<><>>}<<[]{}>{[][]}>>((<()<>>{()()})([[]()]([]{})))}][[",
	"[<[(<(<{(<<{[[{}<>]]}[{{<>{}}{<><>}}[<[][]>{(){}}]]><{<({}<>){{}[]}>(({}[])({}[])>}[{{{}{}}",
	"<((({{[([<<(<(<>())({}())>({()<>}))({<[][]>([]{})}({(){}}[{}{})))>>])]{<[{[<({(){}}{<>})>{{[{}()](",
	"(<<([<<(<[({(<{}()>[<>[]])[({}{}){<>[]}]}(({()[]}<()()>)([[]()]<(){}>)))(<{{<>{}}[{}[]]}([<>{}][(",
	"<({<(([[((([<<()[]>[<>{}]>](<[(){}]<[]{}>)<[[]][[]<>]>))(([<<><>>{[]{}}]))))[({({<[]()><{}()>}<[",
	"{[{{([(({[{[<[[][]]((){})>[[<><>]([][])]]{([[]{}][<>()]){(()[])[[]<>]}}}[<({{}[]}{<>()})[([]())<{}()>",
	"(<(<[<(<({[[((()()){<><>))[(<>())(()())]][[<<>{}>([]())][{<>()}[<>]]]]}[([{[[]<>]}{<[]{}>}]){<<([]{}){{}[]}>>",
	"[([{(<{[({(([<[][]>[{}<>]]<(<><>)>){({()[]}{()[]})({()}(<>[]))})<[[([]())[(){}]]{[()<>]<[]<>>}]((<<>(",
	"<(<{([{<[{<<[[<>[]]]><<[<>{}][{}{}]>>>([({[]()}<<>()>)[<[]()>]]{<{()<>}>[{[]()}{<>[]}]})}({{[(()[])<[]()",
	"({{<{{[<({{([<()[]><{}<>>]<{[][]}[<><>)>)[<(<>{}){{}()}>(<<>{}>[{}[]])]}<{((()<>)[()()])<((){}){{}{}}>",
	"<{[[{{{[(([<({<><>}{<>()})>({<[]{}>[<><>]})]){<[[[{}[]]{()()}]{[<><>)[{}{}]}]><((([][])){[[]",
	"{[{[[[{[<<{{[[[][]]{{}()}][<<>()>{(){}}}}{{<<>{}>[{}[]]}}}[({<<>[]>{[]{}}}<({}<>)[[]()]>)]>({(<<[]<>>",
	"<[[{(<({<({<([()<>])<{<>[]}{{}()}>>})<([<(<>[])<()()>>(<(){}>[{}[]])>([[{}[]]({})]{{<><>}<<>[]>}))[[{[{}",
	"<{{[[{<(<[((((()<>)<{}[]>)<((){})[(){}]>))][[<[({}<>){()}]<[<>{}]{[][]}>>({[()()][()()]})]]>)>}]<<<(",
	"[<[<{(<{<[(([<[][]>])[((<>()>([][]))[{()}{{}<>}]])({[[{}<>](()())](({}())<()[]>)}<{(<>())[(){}]",
	"[<<[[(([[[({[{()[]}(<>{})]{[{}]<<>()>}}{[{[]<>}{()()}]<([][])>})][{((<<>()>(<>[])){{<><>}<()[]>})<[[{}(",
	"<{<([[[([[{[{{[]<>}<{}<>>}](<{()[]}(<>)>{<{}{}>([]{})})}]<<{{{<>{}}{<>()}}<<<>()>[()[]]>}[{[[]",
	"{[[[([{<([(([[()[]]([][])][[<>()](()<>)]]<<{(){}}[(){}]>[{[]()}(<>)]>)<[[[{}()][()]]<({}{}){<>()}>]([[()()]{",
	"(({<(({[(<(({({}[])({}[]]}([[]()]([]())))[<<()<>>{{}()}>[[()()]]])>)<([[[[{}[]]{<>()}]]({([]())(<>())",
	"<[<{({{<{({<{[{}()]{[]}}<[()<>]<(){}>>>[[<[]{}>{{}<>}]<[[]]<<>[]>>]}<{<[<>()][{}{}]>{[()()][<>]}}>)}{{{[[[[]<",
	"(({{([([{{{[{{{}[]}{<><>}}<{{}}(<>())>][(([]{}){<><>})<<[]<>>[<>()]>]}}{{[([<>{}]<()()>)({(){}})]({(<",
	"<<<<(([(([{[<[(){}]([]<>)>]((([][])({}{}))(([]{})<[]>))}])(({{[{{}()}<[]<>>](({}<>)[<>()])}({{[",
	"{{[<[<[([{(({([]())(<>())}[[[]<>]([]{})])[{<[][]>({}{})}{({}[]){()<>}}])}][{{(<(<><>)([]<>)>([<>()](()())))",
	"[[[(<[<([[{[<<(){}>((){})>{<{}<>>({}())}][<<()>([]<>)>([<>[]]<[]>)]}{{<[<>()]<{}<>>><<()<>>{{}<>}>}<<[[]",
	"[<<(<<({<{{(((()())))<[{()()}(()[])][[()()]]>}[{{(()<>)[[]()]}<<<>{}>{<><>}>}[([()[]]<()()>)(",
	"<<{{{<<{{{<[{<[]()>([][])}<{<>{}}{[]{}}>]([{<>()}({}<>)])>}<<[([{}{}][{}])[{{}<>}(()<>)]][<({}())({}()",
	"[<<<{[{(([<<[{<><>}[[]<>]][([][]){[][]}]>[[{()[]}[{}[]]]({()[]}[{}{}]>]>{<[<[]()><<>{}>]<{()[]}([][])>>}]<{",
	"{{([({<[({{{<{[]<>}{<><>}>[<{}{}>{{}()}]}<<<[]()><<>()>}{(()[])[{}{}]}>}}<<([[[]()]{[]<>}]<(()[])([]{})>)[<((",
	"[{<{{<[[<[{<(([]{})<()<>>)(<[][]><[]()])>[{({}<>){{}()}}({()[]}[(){}])]}]>]]{[<({<{<<><>>}<{(",
	"[[<[(<{[<({[(<{}{}>(()()))]<[<[]<>>[{}()]]<({}())([]())>>})>[([[(({}{}){[][]}){{{}<>}({}{})}]]({<[{}[]][()",
	"([(<[<<(<{((<{()[]}([]{})>(<(){}>)){(({}<>){()[]})<[<>[]]<(){}>>}}}>[{{{[[{}[]]]}((<<>()><{}{}>",
	"({{{[(<({[<{<({}<>){()[]}>[<()<>>(<>{})]}<[<{}()>([]<>)]((<>{})[{}[]])>>({[{{}<>}(<><>)}}{({{}()}<<><>>){[",
	"([[<<{<<[<{(<({}<>)<[]<>>>)[[[{}[]][{}<>]]([()[]](<>()>)]}>][{{[[<<><>>][[{}()]<{}[]>]]{<{()[]}{{",
	"{(<<[(<<({[{<[[]()){(){}}>{[[]{}]<{}{}>}}<{<()>[()()]}[<()()><<><>>]>]<<[<[][]><<>{}>]{{[]",
	"[<({[({{{<[<({[]{}}[[][]])>]>[{{{{<>{}}({}()}}(<[][]>{[]()})}([<<><>><{}{}>])}{[<<{}{}>(()[])><{<",
	"{([[<[({{<(({(()<>){<>{}}}{[<>[]](<>[])}))<[<[<><>]>[([][]){{}()}]]{<(<>())>}>>}[({[{<<>{}>({}[])}<{[]{}",
	"(<<[[[[[<<(({([]()){<>()}}{[{}<>]<<>()>})[{(<>{}){[]()}}]]<<[<[]()>[[][]]]>[<[[]<>]([]{})>{",
	"{[[([({<[{<(({{}()}([]())){{[]()}[(){}]}){[<[]<>>[[]<>]]<(()){<>{}}>}>({((()())({}<>))([()()](",
	"[({<{<{[[{{<<{[]}>{<{}[]>[{}()]}>([<[]()>(())])}}<(((({}())(<><>))))(<[{(){}}([]())][[{}[]][()[]]]>[{{{}()}",
	"[[((([[{<(<<<{{}{}}{()()}>{[[]()]{(){}}}><<(<>())>([<>{}]{<>{}})>>)>[{({{([]{}){<>[]}}{(()",
	"{([{([<({[[[([{}{}]<<>()>)((()())<[]<>>)][<[(){}]<[]()>>[<<><>>]]]<<<({}<>)[()()])<<[]<>><()()",
	"{<(<<{({{([<{[()<>]{()<>}}<[(){}]<{}[]>>>](([<[]>([]())]{([])((){})})((([][])(()<>})[{{}}<(){}>])))}<<((<",
	"([{<{{<{[{<{<{<><>}{{}()}><<[]()>(()[])>}{<(<>())[<>[]]>{{[][]}<()()>}}><<{(()[]){{}()}}>>)(({[[{}",
	"{{(<({<[<([{([()<>]([]()))<<[]{}>[[]{}]]}{<({}<>)(<>[])><<[]{}>[[]()]>}][{{((){})<{}[]>}<(<><>",
	"<({<{[[{[({(<[()[]]{[][]}>{<[]()><()()>>)<{<(){}>((){})}(([]{})({}{}))>}[<({[]<>}([]<>))<({}){<>[]}",
	"([[{(((({{{[<<{}{}>{()()}>({[]<>}[[]()])][{[{}<>]<()()>}([<><>]))}}{({(<[][]>[<><>])}[{[<>()]<[]{}",
	"([{{<[(<([{{(([]())({}[]))}{{(<>())[<>{}]}<(()[])([]())>}}{<{<()>{{}{}}}>(<[{}]<{}{}>><{{}{}}([]{})>)}]){<",
	"{((<{[<<(<([(<{}<>>)]<(<[]{}>{[][]}}[[()()][[]<>]]>){{{[[]{}][<>()]}<((){})(<>())>}{(([]{})<[][]>)[([]())",
	"{({[<{<{[{[<{([]<>)({}<>)}<({}())({}[])>>{{{<>{}}([][])}[<[][]>]}]}((([<<>{}>{<>[]}][<{}{}><()()>])[",
	"{({{(({(<<<[(<<><>><<>[]>)((())<{}<>>)]>{<[<()<>><[]>]{{<><>}[{}[]]}>}>><{[{([{}{}]{<><>})[{",
	"(<{[[({(<{[[(<[]()><<><>})<([]()){[][]}>]{((<>{})[()<>])}]}{[{{<[]<>>{[][]}}[<[]<>><<><>>]}]{{{{",
	"<<(({[[{[<((<[{}{}]{[]()}>{{()[]}<<><>>})([[(){}](<>[])]<<{}[]>[(){}]>))<<{{{}[]}({}{})}(<()<>>[",
	"",
}
