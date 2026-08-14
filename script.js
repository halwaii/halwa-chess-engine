var board = null;
var game = new Chess();

// WebSocket connection to Go Server
var ws = new WebSocket("wss://jayyadav.site/ws");

ws.onmessage = function(event) {
    var data = JSON.parse(event.data);
    var engineMove = data.move; // Engine returned a string like "e7e5"

    if (engineMove && engineMove !== "0000") {
        // Convert UCI string to chess.js format
        var sourceSq = engineMove.substring(0, 2);
        var targetSq = engineMove.substring(2, 4);
        var promo = engineMove.length > 4 ? engineMove.charAt(4) : 'q';

        game.move({
            from: sourceSq,
            to: targetSq,
            promotion: promo
        });
        
        // Update board UI
        board.position(game.fen());
    }
};

function onDragStart (source, piece, position, orientation) {
  if (game.game_over() || piece.search(/^b/) !== -1) return false;
}

function onDrop (source, target) {
  var move = game.move({
    from: source,
    to: target,
    promotion: 'q' 
  });

  if (move === null) return 'snapback';
  
  board.position(game.fen());

  // Bhejo current board state (FEN) engine ko!
  if (!game.game_over()) {
      ws.send(JSON.stringify({ fen: game.fen() }));
  }
}

function onSnapEnd () {
  board.position(game.fen());
}

var config = {
  draggable: true,
  position: 'start',
  onDragStart: onDragStart,
  onDrop: onDrop,
  onSnapEnd: onSnapEnd,
  pieceTheme: 'chessboard/img/chesspieces/wikipedia/{piece}.png'
};

board = Chessboard('myBoard', config);