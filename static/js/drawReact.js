window.onload = function () {
	const canvas = document.getElementById("terminalCanvas");
	const ctx = canvas.getContext("2d");
	const moveCursor = document.querySelector(".move-cursor");

	// Set canvas to full page size
	canvas.width = window.innerWidth;
	canvas.height = window.innerHeight;
	canvas.style.position = "absolute";
	canvas.style.top = "0";
	canvas.style.left = "0";
	canvas.style.pointerEvents = "none"; // Don’t block clicks
	console.log(canvas.width, canvas.height);

	let progress = 0; // 0 to 1, tracks drawing

	function drawTerminalRect() {
		ctx.clearRect(0, 0, canvas.width, canvas.height);

		// Position under .move-cursor
		const rect = moveCursor.getBoundingClientRect();
		const x = rect.left;
		const y = rect.bottom + 5; // 5px below text
		const width = rect.width;
		const height = 50; // Fixed height, like a terminal window
		const perimeter = 2 * (width + height);
		const drawnLength = perimeter * progress;

		ctx.strokeStyle = "#ff5555"; // Rosé Pine red-ish (adjust to match your theme)
		ctx.lineWidth = 2;
		ctx.beginPath();

		// Draw like a terminal typing the outline
		if (drawnLength <= width) {
			// Top edge
			ctx.moveTo(x, y);
			ctx.lineTo(x + drawnLength, y);
		} else if (drawnLength <= width + height) {
			// Top + right edge
			ctx.moveTo(x, y);
			ctx.lineTo(x + width, y);
			ctx.lineTo(x + width, y + (drawnLength - width));
		} else if (drawnLength <= 2 * width + height) {
			// Top + right + bottom
			ctx.moveTo(x, y);
			ctx.lineTo(x + width, y);
			ctx.lineTo(x + width, y + height);
			ctx.lineTo(x + width - (drawnLength - (width + height)), y + height);
		} else {
			// Full rectangle
			ctx.moveTo(x, y);
			ctx.lineTo(x + width, y);
			ctx.lineTo(x + width, y + height);
			ctx.lineTo(x, y + height);
			ctx.lineTo(x, y + (height - (drawnLength - (2 * width + height))));
		}

		ctx.stroke();
		progress += 0.03; // Fast, terminal-like speed (~0.5 seconds)
		if (progress <= 1) {
			requestAnimationFrame(drawTerminalRect);
		}
	}

	// Redraw on resize
	window.addEventListener("resize", () => {
		canvas.width = window.innerWidth;
		canvas.height = window.innerHeight;
		progress = 0;
		drawTerminalRect();
	});

	drawTerminalRect();
};
