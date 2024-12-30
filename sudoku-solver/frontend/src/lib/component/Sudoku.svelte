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
				class={`w-10 h-10 text-center text-lg font-bold border
					${(Math.floor(rowIndex / 3) + Math.floor(colIndex / 3)) % 2 === 0 ? 'bg-gray-100' : 'bg-white'}
					${colIndex % 3 === 0 ? 'border-l-2' : ''}
					${rowIndex % 3 === 0 ? 'border-t-2' : ''}
					${colIndex === 8 ? 'border-r-2' : ''}
					${rowIndex === 8 ? 'border-b-2' : ''}
					border-gray-500`}
			/>
		{/each}
	{/each}
</div>
