const fs = require('fs');
const path = require('path');

const targetDir = 'c:\\DEV\\git-repos\\public\\svelte-memorycarddigital-landingpage\\src\\routes\\contact';
const targetFile = path.join(targetDir, '+page.svelte');

console.log('Starting cleanup...');

try {
    if (fs.existsSync(targetFile)) {
        console.log('Deleting file: ' + targetFile);
        fs.unlinkSync(targetFile);
        console.log('File deleted.');
    } else {
        console.log('File not found: ' + targetFile);
    }

    if (fs.existsSync(targetDir)) {
        console.log('Deleting directory: ' + targetDir);
        fs.rmdirSync(targetDir);
        console.log('Directory deleted.');
    } else {
        console.log('Directory not found: ' + targetDir);
    }
} catch (error) {
    console.error('Error during cleanup:', error);
}
