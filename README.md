# Gas Station

> NeetCode 150 · Greedy — two facts carry it: a non-negative total guarantees some start works, and a failed stretch disqualifies all its starts at once. Go, stdlib only, tested.

## Problem

Stations on a circle each give gas and charge a travel cost to the next. From which station can the full loop be driven (unique if it exists), or −1?

## Approach

Fact one: if total gas ≥ total cost, a valid start exists (provable by induction — some station of maximum prefix-deficit works), so feasibility is just the global sum. Fact two kills the O(n²) start-by-start simulation: when the running tank from candidate `start` goes negative at station i, then *every* station in [start, i] fails too — each would arrive at i with no more fuel than the current attempt did (their prefix had non-negative tank up to their own start). So the next candidate is i+1, and the scan never backs up: one pass, candidate monotonically advancing, answer = the last candidate if the total allows any.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
