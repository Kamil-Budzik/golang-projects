<script lang="ts">
	export let board: number[][] = Array(9).fill(Array(9).fill(0));
	export let onInput: (updatedBoard: number[][]) => void = () => {};

	function handleInput(event: Event, row: number, col: number): void {
		const target = event.target as HTMLInputElement;
		const value = parseInt(target.value) || 0;

		if (value >= 0 && value <= 9) {
			const updatedBoard = board.map((r, i) =>
				i === row ? r.map((cell, j) => (j === col ? value : cell)) : r
			);
			onInput(updatedBoard);
		}
	}
</script>

<div class="grid grid-cols-9 gap-1 rounded-lg bg-gray-100 p-4 shadow-md">
	{#each board as row, rowIndex}
		{#each row as cell, colIndex}
			<input
				type="text"
				maxlength="1"
				value={cell === 0 ? '' : cell}
				on:input={(e) => handleInput(e, rowIndex, colIndex)}
			/>
		{/each}
	{/each}
</div>
