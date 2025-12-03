const fs = require('fs');
const path = 'c:\\DEV\\git-repos\\public\\svelte-memorycarddigital-landingpage\\src\\routes\\contact';

try {
    if (fs.existsSync(path)) {
        fs.rmSync(path, { recursive: true, force: true });
        console.log('Successfully deleted ' + path);
    } else {
        console.log('Path does not exist: ' + path);
    }
} catch (err) {
    console.error('Error deleting path:', err);
}
