const fs = require('fs');
const path = require('path');

const targetFile = path.join(__dirname, 'src/routes/contact/+page.svelte');
const logFile = path.join(__dirname, 'delete_log.txt');

try {
    if (fs.existsSync(targetFile)) {
        fs.unlinkSync(targetFile);
        fs.writeFileSync(logFile, 'File deleted successfully');
    } else {
        fs.writeFileSync(logFile, 'File not found');
    }
} catch (e) {
    fs.writeFileSync(logFile, 'Error: ' + e.message);
}
