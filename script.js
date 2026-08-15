var board = null;
var game = new Chess();
var playerColor = 'w'; // Default player color

// WebSocket connection
var ws = new WebSocket("wss://jayyadav.site/ws"); 

ws.onmessage = function(event) {
    var data = JSON.parse(event.data);
    
    if (data.visits !== undefined) {
        document.getElementById("visitCount").innerText = data.visits;
        document.getElementById("gameCount").innerText = data.games;
        document.getElementById("winCount").innerText = data.engineWins;
        document.getElementById("drawCount").innerText = data.draws;
        return;
    }

    var engineMove = data.move; 

    if (engineMove && engineMove !== "0000") {
        var sourceSq = engineMove.substring(0, 2);
        var targetSq = engineMove.substring(2, 4);
        var promo = engineMove.length > 4 ? engineMove.charAt(4) : 'q';

        game.move({ from: sourceSq, to: targetSq, promotion: promo });
        board.position(game.fen());
        checkGameOver();
    }
};

function onDragStart (source, piece, position, orientation) {
  
  if (game.game_over() || piece.charAt(0) !== playerColor) return false;
}

function onDrop (source, target) {
  var move = game.move({ from: source, to: target, promotion: 'q' });

  if (move === null) return 'snapback';
  board.position(game.fen());

  checkGameOver();

  if (!game.game_over()) {
      ws.send(JSON.stringify({ fen: game.fen() }));
  }
}

function onSnapEnd () { board.position(game.fen()); }

function checkGameOver() {
    if (game.game_over()) {
        if (game.in_draw() || game.in_stalemate()) {
            ws.send(JSON.stringify({ result: "draw" }));
            alert("Game Drawn! ");
        } else {
            
            let winner = game.turn() === 'w' ? 'Black' : 'White';
            if (winner.charAt(0).toLowerCase() !== playerColor) {
                ws.send(JSON.stringify({ result: "enginewin" }));
            }
            alert("Checkmate! " + winner + " wins! ");
        }
    }
}

function startNewGame(color) {
    game.reset();
    playerColor = color === 'white' ? 'w' : 'b';
    board.orientation(color);
    board.position('start');
    
    if (playerColor === 'b') {
        ws.send(JSON.stringify({ fen: game.fen() }));
    }
}


function downloadPGN() {
    if (game.pgn() === "") {
        alert("first play game");
        return;
    }
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(game.pgn()));
    element.setAttribute('download', 'HalwaEngine_Game.pgn');
    element.style.display = 'none';
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}

var whiteSquareGrey = '#a9a9a9';
var blackSquareGrey = '#696969';

function removeGreySquares() {
    $('#myBoard .square-55d63').css('background', '');
}

function greySquare(square) {
    var $square = $('#myBoard .square-' + square);
    var background = whiteSquareGrey;
    if ($square.hasClass('black-3c85d')) {
        background = blackSquareGrey;
    }
    $square.css('background', background);
}

function onMouseoverSquare(square, piece) {
    
    var moves = game.moves({
        square: square,
        verbose: true
    });
    if (moves.length === 0) return;

    // source square highlight
    greySquare(square);

    // highlight every legal move
    for (var i = 0; i < moves.length; i++) {
        greySquare(moves[i].to);
    }
}

function onMouseoutSquare(square, piece) {
    removeGreySquares();
}
var config = {
  draggable: true,
  position: 'start',
  onDragStart: onDragStart,
  onDrop: onDrop,
  onMouseoverSquare: onMouseoverSquare,   
  onMouseoutSquare: onMouseoutSquare,     
  onSnapEnd: onSnapEnd,
  pieceTheme: 'chessboard/img/chesspieces/wikipedia/{piece}.png'
};

board = Chessboard('myBoard', config);