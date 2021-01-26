:- use_module(library(clpb)).
circuit(A,B,C) :- sat((A*B*C) + (A * (~B + ~C))), labeling([A,B,C]).
