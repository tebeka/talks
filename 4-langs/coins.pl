coin(dollar, 100).
coin(half, 50).
coin(quarter, 25).
coin(dime,10).
coin(nickel,5).
coin(penny,1).

change(0, []).
change(A, [(C, Num)|L]) :-
    coin(C, V),
    A >= V,
    Num is floor(A / V),
    B is A mod V,
    change(B, L).


% vim: ft=prolog
