mountain(everest).
mountain(aconcagua).
mountain(mckinley).
mountain(kilimanjaro).
mountain(elbrus).
mountain(vinson).
mountain(puncakjaya).
location(everest,asia).
location(aconcagua,southamerica).
location(mckinley,northamerica).
location(kilimanjaro,africa).
location(elbrus,europe).
location(vinson,antarctica).
location(puncakjaya,australia).
height(everest,29029).
height(aconcagua,22841).
height(mckinley,20312).
height(kilimanjaro,19340).
height(elbrus,18510).
height(vinson,16050).
height(puncakjaya,16023).
highest(everest).
climber(john).
climber(kelly).
climber(maria).
climber(derek).
climber(thyago).
certified(john).
certified(kelly).
certified(maria).
certified(derek).
not(certified(thyago)).

higher(X,Y) :-
    height(X,H),
    (H>Y).

climb(john,X) :-
    mountain(X),
    certified(john),
    location(X,northamerica).

climb(kelly,X) :-
    mountain(X),
    certified(kelly),
    height(X,Y),
    Y >= 20000.

climb(maria,X) :-
    mountain(X),
    certified(maria).

climb(derek,X) :-
    mountain(X),
    certified(john),
    (location(X,europe);
    location(X,africa);
    location(X,australia)),
    height(X,Y),
    Y =< 19000.

climb(thyago) :-
    false.
