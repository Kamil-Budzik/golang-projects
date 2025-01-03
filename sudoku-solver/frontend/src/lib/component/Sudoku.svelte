<script lang="ts">
	export let board: number[][] = Array(9).fill(Array(9).fill(0));

	function handleInput(event: Event, row: number, col: number): void {
		const target = event.target as HTMLInputElement;
		const value = parseInt(target.value) || 0;

		if (value >= 0 && value <= 9) {
			board[row][col] = value;
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
				oninput={(e) => handleInput(e, rowIndex, colIndex)}
				class={`h-10 w-10 border text-center text-lg font-bold ${(Math.floor(rowIndex / 3) + Math.floor(colIndex / 3)) % 2 === 0 ? 'bg-gray-100' : 'bg-white'} ${colIndex % 3 === 0 ? 'border-l-2' : ''} ${rowIndex % 3 === 0 ? 'border-t-2' : ''} ${colIndex === 8 ? 'border-r-2' : ''} ${rowIndex === 8 ? 'border-b-2' : ''} border-gray-500`}
			/>
		{/each}
	{/each}
</div>
