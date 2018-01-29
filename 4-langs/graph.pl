edge(a,e).
edge(e,d).
edge(d,c).
edge(c,b).
edge(e,c).
edge(f,b).

path(X,X).
path(X,Y) :- edge(X, Z), path(Z, Y).
