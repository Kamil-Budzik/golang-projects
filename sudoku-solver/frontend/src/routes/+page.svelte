<script lang="ts">
	import Sudoku from '$lib/component/Sudoku.svelte';
	import axios from 'axios';

	let board: number[][] = Array(9).fill(null).map(() => Array(9).fill(0));

	async function solveSudoku() {
		try {
			const apiURL = import.meta.env.VITE_API_BASE_URL;
			const res = await axios.post(apiURL + '/solve', { board });

			console.log(res.data);

			// const solvedBoard = await response.json();
			// board = solvedBoard.board; // Assuming the API returns the solved board in this format
			// console.log('Solved board:', solvedBoard);
		} catch (error) {
			console.error('Failed to solve Sudoku:', error);
		}
	}
</script>

<main class="grid min-h-screen w-full place-items-center">
	<Sudoku bind:board />
	<button
		on:click={solveSudoku}
		class="mt-4 rounded bg-blue-500 px-4 py-2 text-white shadow-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300"
	>
		Solve
	</button>
</main>
