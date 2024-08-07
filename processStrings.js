function processStrings(inputArray) {
    // Gabungkan menjadi satu string
    let combinedString = inputArray.join('');

    // Inisialisasi objek untuk menghitung frekuensi huruf
    let frequency = {};

    // Hitung frekuensi huruf
    for (let char of combinedString) {
        if (!frequency[char]) frequency[char] = 0;
        frequency[char]++;
    }

    // Buat array dari karakter dan frekuensinya
    let charArray = Object.keys(frequency).map(char => {
        return { char, freq: frequency[char] };
    });

    // Urutkan berdasarkan frekuensi dan aturan tambahan
    charArray.sort((a, b) => {
        if (b.freq === a.freq) {
            if (a.char.toLowerCase() === b.char.toLowerCase()) {
                return a.char > b.char ? 1 : -1;
            }
            if (a.char === a.char.toUpperCase() && b.char === b.char.toLowerCase()) {
                return -1;
            }
            if (a.char === a.char.toLowerCase() && b.char === b.char.toUpperCase()) {
                return 1;
            }
            return a.char.toLowerCase() > b.char.toLowerCase() ? 1 : -1;
        }
        return b.freq - a.freq;
    });

    // Buat string hasil akhir dengan hanya satu kali kemunculan setiap karakter
    let result = '';
    let seen = new Set();

    for (let item of charArray) {
        if (!seen.has(item.char)) {
            result += item.char;
            seen.add(item.char);
        }
    }

    return result;
}

// cara penggunaan
let inputArray = ["Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"];
let output = processStrings(inputArray);
console.log(output); 
