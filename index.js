if (process.env['HELLO'] === 'WORLD') {
    console.log(`Hello, ${process.env['HELLO']}`);
} else {
    process.exitCode = 1;
    console.error(`HELLO env mismatch: "${process.env['HELLO']}"`);
}
