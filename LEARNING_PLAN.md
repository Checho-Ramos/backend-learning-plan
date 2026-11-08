# Backend Engineering Plan — Student Edition (AI-Assisted)

**Who this is for:** you know programming basics (variables, loops, functions, conditionals, basic Git) and want to become a junior backend developer.
**Goal:** build, test, deploy, and explain a real service.
**Duration:** 16 weeks, 12–16 h/week.
**Stack:** Go, PostgreSQL, Redis, one document database, Docker.
**You will use AI.** This plan tells you exactly how, so it accelerates you instead of replacing you.

---

## 1. Setup (do this before week 1)

- Go toolchain, Git, Docker, an editor, `psql`, `redis-cli`, `curl`.
- A GitHub account and one repo per phase. Commit daily. Every change lands through a pull request: branch, commit, PR, CI, merge.
- Access to an AI assistant and, if possible, a coding agent that can read the repo.
- Copy this document into the repo as `LEARNING_PLAN.md`.
- Create `PROJECT.md`, `COACH.md`, and `progress.md` at the repo root (templates in section 2). `PROJECT.md` is the brief you give the AI every session:

```markdown
# Project brief
Stack: Go 1.2x, PostgreSQL, Redis, Docker Compose.
Commands: `go test ./...`, `go test -race ./...`, `go vet ./...`, `staticcheck ./...`, `docker compose up --build`.
Layout: cmd/api, internal/<domain>, migrations, docs/adr, notes.
Rules: stdlib first; no ORM; SQL by hand; wrap errors; context everywhere; table-driven tests.
Definition of done: runs from clean checkout, tests fail when logic is broken, README updated.
Never: add a dependency without asking, rewrite files I did not ask about, or write code I cannot explain.
```

**Repo layout**

```
LEARNING_PLAN.md    this document
PROJECT.md          project brief
COACH.md            AI role and rules
progress.md         current state, updated every session
cmd/api/            main package
internal/<domain>/  business logic, one package per domain
migrations/         SQL migrations
docs/adr/           architecture decision records
notes/              weekly AI log, tradeoffs, postmortems
```

**Week 0 starter task (do this yourself)**

1. Create a public repo with a `README.md` and a `main` branch.
2. Add one Go function, for example `Greet(name string) string`, plus a table-driven test.
3. Add a GitHub Actions workflow that runs `go test ./...` on every push and pull request.
4. Create a branch, change the function's behavior, and update the test to match.
5. Open a pull request into `main`. Describe what changed and why. Wait for CI to pass.
6. Review your own diff, merge, and delete the branch.

---

## 2. The AI contract

AI is your pair programmer and tutor. It is not your autopilot.

1. **Attempt first.** Spend 20 minutes on the problem before asking. Write down what you tried.
2. **Ask for hints, not answers.** Use the hint ladder: concept, then API, then pseudocode, then code. Stop as early as you can.
3. **Never accept code you cannot explain.** After any AI help, close the chat and explain the code aloud. If you cannot, delete it and redo it.
4. **One concept per diff.** No 300-line AI dumps. Small steps, run tests after each.
5. **You write the tests first.** AI may suggest edge cases. You write the failing test, then make it pass.
6. **You own the core logic.** AI may do boilerplate, docs, test scaffolding, and config. Core business logic: you write the first version.
7. **AI reviews, you decide.** Ask it to critique, not rewrite. Rejected suggestions go in your notes with a reason.
8. **Cold check weekly.** 30 minutes with AI off: implement one small piece, then ask AI to review it.
9. **Three-strike rule.** If AI gives you the answer three times on the same concept, stop. Do 30 minutes of reps on that concept, then redo it cold.
10. **Log it.** `notes/wkNN-ai-log.md`: what you asked, what it suggested, what you kept, what you rejected and why.

**AI red flags:** chat longer than code; accepting code you cannot explain; no failing test; AI writing your ADR or postmortem; AI doing the boss check; rewriting instead of understanding.

### 2.1 `COACH.md` — the AI's role

Put this in the repo root and load it at the start of every session. It turns a general assistant into a coach that follows this plan.

```markdown
# Coach brief

You are my backend coach for the 16-week plan in LEARNING_PLAN.md.
Read LEARNING_PLAN.md, progress.md, and PROJECT.md before every answer.

## Role
- Teach, do not do. I write the code.
- One step at a time. Ask one question at a time.
- Do not write implementation code unless I explicitly say "show me the answer"
  and I have already attempted it for 20 minutes.
- Default to hints: concept, then API, then pseudocode. Stop at the smallest
  hint that unblocks me.
- When I ask for review, critique my diff. Do not rewrite it.
- Grade against the mission's pass condition. Say pass or fail and cite evidence.
- Propose updates to progress.md as a diff. I apply them.
- Never mark work done that I have not demonstrated.
- Never advance the week because time passed. Advance only when the pass
  condition is met.
- If I ask you to "just do it", refuse once and offer the smallest next step.
- Quiz me weekly. Grade honestly.
- Keep answers under 15 lines unless I ask for depth.

## Session start
1. Read LEARNING_PLAN.md, progress.md, and PROJECT.md.
2. Tell me: current week, mission, next step, and one question that tests my
   understanding.
3. Wait for my answer. Do not write code.

## Red flags
If I paste lots of generated code, cannot explain it, or skip tests, call it out
and stop the session.
```

### 2.2 `progress.md` — the AI's memory

`progress.md` is the single source of truth for where you are. Update it at the end of every session and commit it. Without it, the AI starts cold and mis-scopes the work.

```markdown
# Progress

Updated: 2026-01-01
Week: 0
Mission: Boot the Rig
XP: 0
Level: Initiate
Streak: 0
Badges: none

## This mission's pass condition
- [ ] ...
- [ ] ...
- [ ] ...

## Last session
- Did:
- Learned:
- Blocked by:
- AI used for:

## Next step (one thing)
...

## Gate log
| Week | Gate | Result | Evidence |
|---|---|---|---|
| | | | |

## Notes
- ...
```

---

## 3. The loop

**Daily (2–3 h)**

| Time | Activity |
|---|---|
| 5 min | Open `TODO`, pick the one thing |
| 20–30 min | Interactive rep (toolbox below) |
| 20 min | Attempt the build with no AI |
| 60–90 min | Build, using hints only when stuck |
| 10 min | Run tests, commit, write tomorrow's one-liner |
| 5 min | Update the AI log if you used it |

**Weekly**

- Mon: 30-min briefing. Read the mission, write 3 acceptance checks, paste `PROJECT.md` to the AI and ask for a plan only.
- Tue–Thu: build blocks.
- Fri: tests, reps, cold check.
- Sat: boss check (timer, AI closed), demo, update progress board.
- Sun: off.

**Weekly AI workflow**

1. **Start:** load `COACH.md`, `progress.md`, `PROJECT.md`, and the mission. Ask for the next step only. No code.
2. **Plan:** ask it to ask you questions and produce a step list, each step under 30 minutes and ending in something runnable.
3. **Build:** you code. When stuck, use the hint ladder.
4. **Review:** after it runs, ask for a review of the diff. Fix what you agree with.
5. **Quiz:** ask it to quiz you on the week's topic, then grade your answers.
6. **Close:** update `progress.md` and `notes/wkNN-ai-log.md`, commit, and state the single next step.

---

## 4. Prompt library (copy-paste)

**Session start (coach mode)**
> Read COACH.md, progress.md, and PROJECT.md. Tell me the current week, the mission, today's next step, and one question to test my understanding. Do not write code. Wait for my answer.

**Gate check**
> Using the week [N] pass condition and the evidence below, grade me pass or fail. Cite what is missing. Do not be generous. Evidence: [test output, demo notes, diff].

**Session close**
> Here is what I did, what I learned, and what is blocked. Propose an update to progress.md: XP, badges, pass-condition checkboxes, and the single next step. Do not mark anything done I have not demonstrated.

**Plan, no code**
> I am implementing [mission]. Read PROJECT.md. Ask me 5 questions before suggesting a design. Do not write code yet. Then give me a step list where each step is under 30 minutes and ends with something runnable.

**Hint ladder**
> I am stuck on [problem]. I tried [X] and got [error]. Give me the smallest possible hint, not the solution. Level 1: concept. Level 2: which API or tool. Level 3: pseudocode. Wait for me to ask before going to the next level.

**Debug, no fix**
> Here is the failing test, the error, and the relevant code. Ask me for any missing information, then list 3 hypotheses ranked by likelihood. Do not edit the code. I will test them.

**Senior review**
> Review this diff as a senior Go reviewer. Focus on correctness, error handling, context usage, concurrency, SQL correctness, and test quality. Point to specific lines. Do not rewrite the code. End with the top 3 things to fix.

**Quiz me**
> Quiz me on [topic] with 5 questions, easiest to hardest. One question at a time. Wait for my answer, then grade it and explain what I missed. Do not reveal the answer before I try.

**Scope cutter**
> I have 90 minutes. Cut this mission to the smallest version that still demonstrates [concept]. List what to skip and what to keep. No code.

**Requirement generator**
> Act as a demanding client for [capstone]. Write 8 acceptance criteria, including 2 failure cases. Do not design the solution.

**Postmortem**
> Help me write a blameless postmortem. Ask for the timeline, impact, detection, and root cause before writing anything. Do not invent facts.

**Explain-back**
> I will explain this code to you. Tell me where my explanation is wrong or incomplete. Do not correct me until I finish.

---

## 5. Progression

**XP**

| Action | XP |
|---|---|
| Daily rep | 10 |
| Mission shipped | 50 |
| Boss check passed | 100 |
| Phase gate passed | 100 |
| Capstone milestone | 200 |
| Demo to someone outside the course | 50 |

**Levels:** Initiate 0 · Apprentice 300 · Builder 700 · Systems Smith 1200 · Operator 1800 · Shipper 2500.

**Streak:** 5 sessions/week. One freeze token per week.

**Badges**
- **Race-Free** — `-race` clean under load.
- **Ghost in the Machine** — prove and fix a goroutine leak.
- **Index Whisperer** — 10x query improvement with plans documented.
- **Migration Monk** — forward and rollback, zero data loss.
- **Cache Cooler** — survives a cache stampede.
- **Idempotent One** — duplicate POST returns the same result, once.
- **No Lost Job** — worker killed mid-job, no loss or duplication.
- **Chaos Tamer** — passes the failure run.
- **One-Command** — fresh clone to running in one command.
- **Stranger Test** — an outsider runs it from the README.
- **On-Call** — fixes an injected incident in 30 minutes.

---

## 6. Toolbox (core only)

| Skill | Resource | Use |
|---|---|---|
| Go | [A Tour of Go](https://go.dev/tour/) | Syntax reps in browser |
| Go | [Go by Example](https://gobyexample.com/) | Run, break, modify examples |
| Go | [Exercism Go](https://exercism.org/tracks/go) | 20–30 min reps, free mentoring |
| Go | [Go Interview Practice](https://app.gointerview.dev/) | Timed challenges with tests |
| Concurrency | [Gist of Go](https://antonz.org/go-concurrency/) | Interactive goroutine/channel chapters |
| SQL | [SQLBolt](https://sqlbolt.com/) | First pass |
| SQL | [PGExercises](https://pgexercises.com/) | Postgres depth |
| SQL | [SQLRide](https://sqlride.com/) | No-signup drills |
| SQL | [sqlpractice.in](https://www.sqlpractice.in/practice) | Scenario problems |
| Redis | [Redis Sandbox](https://redis.io/try/sandbox/) | Command reps |
| MongoDB | [MongoDB Playground](https://mongoplayground.net/) | Aggregation reps |
| Systems | [SystemBlocks](https://systemblocks.dev/) | Scaling/caching/queue sims |
| Containers | [iximiuz Labs](https://labs.iximiuz.com/) | Docker/Linux VMs |
| Containers | [Killercoda](https://killercoda.com/) | Browser scenarios |
| HTTP | [Learn API](https://learn-api.com/) | Request/response reps |
| HTTP | [API Playground](https://playgroundapi.com/playground) | Inspect real responses |
| Git | [Learn Git Branching](https://learngitbranching.js.org/) | Branching reps |

Optional after a phase gate: CodeCrafters, Boot.dev, HackerRank SQL, DB Fiddle, Neon/Supabase.

---

## 7. Missions

Reps are 20–30 minutes. Boss checks are timed and done with AI closed; AI reviews afterwards.

| Week | Mission | Reps | Boss check — pass when | XP |
|---|---|---|---|---|
| 0 | Boot the Rig | Git branching, 3 levels; Tour of Go, 1 module | Repo has one Go function + test; a branch updates it, PR is green, reviewed, merged | 50 |
| 1 | Data Wrangler: CSV CLI | Tour of Go basics; 3 easy Exercism | Streams 1 GB under 200 MB RAM; tests fail if logic breaks | 50 |
| 2 | API Awakens: JSON HTTP | Go by Example HTTP/JSON; 2 medium Exercism | New endpoint in 25 min, tests included, explained cold | 50 |
| 3 | Concurrency Arena | Gist of Go goroutines/channels; 2 pipeline exercises | 10k jobs, clean cancel, `-race` clean · Race-Free | 100 |
| 4 | Refactor Raid | slog/embed examples; 3 Go challenges | Feature request shipped in under an hour, at most 2 files touched · Go gate | 100 |
| 5 | Schema Forge | SQLBolt; PGExercises basic | Constraints reject 5 kinds of bad data, no app checks | 50 |
| 6 | Index Hunter | PGExercises aggregates/window; 5 SQLRide | 5 queries 10x faster; plans in `notes/` · Index Whisperer | 100 |
| 7 | Persistence Run | 5 sqlpractice.in; DB Fiddle | Forward and rollback migrations, zero data loss · SQL gate | 100 |
| 8 | Cache & Chaos | 10 Redis commands; 3 Mongo aggregations | Stampede test passes; Redis down degrades, not crashes · Cache Cooler · NoSQL gate | 100 |
| 9 | Contract & Keys | Learn API auth/idempotency; API Playground | Duplicate POST once; authz matrix passes; spec matches · Idempotent One · API gate | 100 |
| 10 | Boundary Wars | 2 SystemBlocks modules | Swap module implementation, handlers untouched, no internal imports | 50 |
| 11 | Async Engine | SystemBlocks queues/retries | Worker killed mid-job; no loss, no duplicate · No Lost Job · Architecture gate | 100 |
| 12 | Failure Lab | SystemBlocks CAP/partitioning | Random 500s, latency, restarts: still correct; 3 failure modes written · Chaos Tamer | 100 |
| 13 | Ship It | 3 iximiuz Docker labs; Killercoda | Fresh clone to running in one command; CI catches a planted bug · One-Command | 100 |
| 14 | Pressure Test | Killercoda Linux/network | p95 under 200 ms at target load; one real security hole fixed | 100 |
| 15 | Capstone Sprint | Docs only | Deployed and demoed to a stranger · Stranger Test | 200 |
| 16 | Operation Day | None | Incident diagnosed and fixed in 30 min from logs, metrics, runbook; postmortem · On-Call · Capstone gate | 200 |

**AI mode by phase**

- **Weeks 0–4:** AI off for the first 20 minutes; then hints only. No generated Go logic.
- **Weeks 5–8:** AI may explain plans and review SQL. You write every query.
- **Weeks 9–12:** AI as reviewer only; it may suggest tests, not implementations.
- **Weeks 13–16:** AI may write Docker and CI boilerplate. You own config, secrets, and the runbook.

---

## 8. Definition of done (every mission)

- Runs from a clean checkout using the README.
- Tests cover the win condition and fail if the logic is inverted.
- `go vet`, `staticcheck`, and `-race` are clean where applicable.
- Code is small; no dependency added without a note.
- `notes/wkNN.md`: 3-sentence explanation, 1 tradeoff, 1 thing that broke.
- `notes/wkNN-ai-log.md`: 5 lines on AI use.
- `progress.md` updated and committed.
- One commit history that shows small steps, not one giant dump.

---

## 9. Boss checks without a mentor

1. Write the acceptance checks yourself on Monday from the mission's win condition.
2. On Saturday, set a timer, close the AI, and do the boss check.
3. Record a 3-minute demo: run it, show the test failing then passing, explain the tradeoff.
4. Only then ask AI to review. Compare its review with your own notes.
5. If you fail, do not move on. Spend the next week on that phase at smaller scope.
6. Every two weeks, show the demo to a human. A friend is enough. Feedback is the point.

**Phase gates:** Go (wk 4), SQL (wk 7), NoSQL (wk 8), API (wk 9), Architecture (wk 11), Capstone (wk 16). Passing requires the demo, the tests, and a cold explanation.

---

## 10. Capstone (weeks 15–16)

Pick a domain you actually like. Scope is fixed.

- Go service, PostgreSQL, Redis.
- Auth plus resource-level authz.
- One async worker with retry/backoff and dead-letter handling.
- OpenAPI spec.
- Docker Compose, one command to start.
- CI: tests, race detector, static analysis, migrations.
- README, runbook, one ADR per major decision.
- Postmortem of one real bug.

Out of scope: Kubernetes, service mesh, event sourcing, multi-region, Kafka.

**Use AI here as a client and reviewer, not a builder.** It writes acceptance criteria and reviews your design. You write the service.

---

## 11. Reference shelf (not curriculum)

Read only when blocked or curious; 20 minutes max.

- *Effective Go* — reference.
- *Learning Go*, Bodner — lookup.
- *100 Go Mistakes* — one per day if interested.
- *Use The Index, Luke* — when a query is slow.
- *A Philosophy of Software Design* — one chapter per phase.
- *Designing Data-Intensive Applications* — chapters 1–3 and 5–7, after week 10.
- *Fundamentals of Software Architecture* — reference.
- [The Twelve-Factor App](https://12factor.net/) — 15 minutes, week 11.
- Google SRE Book, chapters 1–3 — week 13.
- [System Design Primer](https://github.com/donnemartin/system-design-primer) — lookup.

---

## 12. Progress board

`progress.md` is the source of truth. This board is a mirror; update both at the end of each week.

**Phase 1 — Go**
- [ ] Boot the Rig
- [ ] Data Wrangler
- [ ] API Awakens
- [ ] Concurrency Arena — Race-Free
- [ ] Refactor Raid — Go gate

**Phase 2 — Data**
- [ ] Schema Forge
- [ ] Index Hunter — Index Whisperer
- [ ] Persistence Run — SQL gate
- [ ] Cache & Chaos — Cache Cooler, NoSQL gate

**Phase 3 — Design**
- [ ] Contract & Keys — Idempotent One, API gate
- [ ] Boundary Wars
- [ ] Async Engine — No Lost Job, Architecture gate
- [ ] Failure Lab — Chaos Tamer

**Phase 4 — Ship**
- [ ] Ship It — One-Command
- [ ] Pressure Test
- [ ] Capstone Sprint — Stranger Test
- [ ] Operation Day — On-Call, Capstone gate

**Score:** _____ XP · Level: _____ · Streak: _____

---

## 13. Boredom protocol

1. **Switch the skin, not the stack.** Same mission, new domain.
2. **Shrink it.** Ship the dumb version, then improve it.
3. **Add a constraint.** No frameworks, no ORM, 100-line limit, single file.
4. **Make it visible.** Deploy, demo, get one piece of feedback.
5. **Raise the stakes.** Move the boss check earlier or add a timer.
6. **Break it to learn it.** Remove the index, kill the cache, drop the connection.

**Behind schedule? Cut scope, not sequence:** drop the document DB, drop tracing, pick the smaller capstone. Never drop tests, migrations, authz, or the postmortem.

---

## 14. Weekly review (fill this in)

```
Week: __  Mission: __  XP earned: __
Shipped: yes/no
Tests that fail when logic breaks: yes/no
AI used for: plan / hint / review / boilerplate / none
Code I could not explain and deleted: 
Tradeoff I chose:
What broke and how I found it:
progress.md committed: yes/no
Next week's one thing:
```
