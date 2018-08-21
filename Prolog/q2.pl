:- dynamic(parent/0).

ancestor(X,Y) :-
    parent(X,Y).
ancestor(X,Y) :-
    parent(Z,Y),
    ancestor(X,Z).

lca(X,Y) :-
    X==Y -> write(X);
    ancestor(X,Y) -> write(X);
    parent(Z,X),lca(Z,Y).

forloop(1).
forloop(N1) :-
    readln([X,Y]),
    assert(parent(X,Y)),
    write('node '),write(X),write(' is the parent node of node '),write(Y),nl,
    M1 is N1-1,
    forloop(M1).

forloop1(0).
forloop1(N2) :-
    readln([X,Y]),
    write('Which node is the LCA of node '),write(X),write(' and node '),write(Y),write('?'),nl,
    lca(X,Y),nl,
    M2 is N2-1,
    forloop1(M2).

:-
    write('Enter the number of nodes > '),
    readln(INPUT1),write('# of nodes'),nl,
    forloop(INPUT1).
:-
    write('Enter the number of queries > '),
    readln(INPUT2),write('# of queries'),nl,
    forloop1(INPUT2).
