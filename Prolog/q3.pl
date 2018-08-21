:- dynamic(edge/0).

bi_edge(X,Y) :- 
    edge(X,Y);
    edge(Y,X).


con(X,Y) :-
    con(X,Y,[]).
con(X,Y, _) :-
    bi_edge(X,Y).
con(X,Y,V) :-
    \+ member(X,V),
    bi_edge(X,Z),
    con(Z,Y,[X|V]).

forloop1(0).
forloop1(N1) :-
    readln([X,Y]),
    assert(edge(X,Y)),
    write('node '),write(X),write(' and node '),write(Y),write(' are connected'),nl,
    M1 is N1-1,
    forloop1(M1).

forloop2(0).
forloop2(N2) :-
    readln([X,Y]),
    write('Are node '),write(X),write(' and node '),write(Y),write(' connected?'),nl,
    (
        \+ con(X,Y) ->
        write('No , node '),write(X),write(' and node '),write(Y),write(' are NOT connected'),nl
    ;
        write('Yes , node '),write(X),write(' and node '),write(Y),write(' are connected'),nl
    ),
    M2 is N2-1,
    forloop2(M2).

:-
    write('Enter the number of nodes and number of edges > '),
    readln([X,Y]),write('# of nodes and # of edges'),nl,
    forloop1(Y).
:-
    write('Enter the number of queries > '),
    readln(Z),write('# of queries'),nl,
    forloop2(Z).
