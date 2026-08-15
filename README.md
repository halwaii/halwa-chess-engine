# Halwa Chess Engine

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![WebSockets](https://img.shields.io/badge/WebSockets-Enabled-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Live-success?style=for-the-badge)

Welcome to **Halwa**, a high-performance, custom-built Chess Engine written entirely from scratch in Go (Golang). It features a blazing-fast backend and a sleek, interactive web interface. 

The engine currently plays at an estimated **~1900 ELO**, but it holds the proud badge of having defeated a 2100-rated Stockfish in a glorious tactical brawl! 

 **Play against Halwa here:** [jayyadav.site](https://jayyadav.site)

---

##  Features

*   **Play as White or Black:** Choose your side or let the engine surprise you. The board dynamically flips based on your choice.
*   **Legal Move Highlights:** Hover over any piece to instantly see all valid moves and captures.
*   **Live Global Stats:** Real-time tracking of Total Visits, Games Played, Engine Wins, and Draws across all players.
*   **PGN Downloads:** Download the complete move history of your game with a single click to analyze it later.
*   **Zero-Lag WebSockets:** Seamless, bi-directional communication between the Go engine and your browser.
*   **Concurrency Safe:** Implements robust mutex locking to handle multiple players simultaneously without crashing the engine.

---

##  Under the Hood

Halwa isn't just a basic move generator; it implements several advanced chess programming techniques to search deep and fast:

*   **Move Ordering & Heuristics:** Prioritizes captures and killer moves to optimize the search tree[cite: 1, 2].
*   **Memoization:** Utilizes Zobrist Hashing and Transposition Tables (TT) to remember previously evaluated positions and save time[cite: 1, 3].
*   **Quiescence Search:** Evaluates active captures at the end of the main search to completely eliminate the Horizon Effect[cite: 1, 3].
*   **Advanced Evaluation:** Calculates king safety, pawn structures, development rewards, and endgame mop-up patterns[cite: 2, 3].

---

##  Run It Locally

Want to test the engine on your own machine? It takes less than a minute:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/jay-xlr8/halwa-chess-engine.git
   cd halwa-chess-engine

2. run the go server
   go run .

3. Play:
Open your browser and navigate to http://localhost:8080.

Halwa is an ongoing journey of learning and optimization. Anyone is absolutely free to contribute!

Whether you want to improve the evaluation function, optimize the search depth, add opening books, or tweak the frontend UI—your PRs are more than welcome.

Fork the Project

Create your Feature Branch (git checkout -b feature/AmazingFeature)

Commit your Changes (git commit -m 'Add some AmazingFeature')

Push to the Branch (git push origin feature/AmazingFeature)

Open a Pull Request

Let's build something awesome together! 🌟
