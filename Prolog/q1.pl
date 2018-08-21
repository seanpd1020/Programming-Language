divisible(X,Y) :- 
    X mod Y =:= 0.
divisible(X,Y) :- 
    Y * Y < X,
    divisible(X,Y + 2).

prime(2).
prime(3).
prime(X) :- 
    Y is 3, X > 3,
    X mod 2 =:= 1,
    \+divisible(X,Y).

nextprime(1,2).
nextprime(2,3).
nextprime(X,Y) :-
    Y is X + 2,
    prime(Y), !.
nextprime(X,Y) :- 
    Z is X + 2,
    nextprime(Z,Y).

goldbach(4,[2,2]).
goldbach(X,Y) :-
    X mod 2 =:= 0,
    X > 4,
    goldbach(X,Y,3).
goldbach(X,[Z,Y],Z) :-
    Y is X - Z,
    prime(Y), !.
goldbach(X,Y,Z) :- Z < X,
    nextprime(Z,Z1),
    goldbach(X,Y,Z1).

:-
    write('Input:'),
    readln(Input),
    goldbach(Input,Output),
    write('Output:'),
    write(Output),nl.
