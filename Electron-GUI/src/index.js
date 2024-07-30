const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require('electron-squirrel-startup')) {
  app.quit();
}

let mainWindow, loginWindow;

const createLoginWindow = () => {
  loginWindow = new BrowserWindow({
    width: 600,
    height: 500,
    resizable: false,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
    },
  });

  loginWindow.loadFile(path.join(__dirname, 'login.html'));

  // Handle closing the login window
  loginWindow.on('closed', () => {
    loginWindow = null;
  });


};

const createMainWindow = () => {
  mainWindow = new BrowserWindow({
    width: 600,
    height: 500,
    resizable:false,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
    },
  });

  mainWindow.loadFile(path.join(__dirname, 'app.html'));

  // Handle closing the main window
  mainWindow.on('closed', () => {
    mainWindow = null;
  });
};

// This method will be called when Electron has finished initialization
app.whenReady().then(() => {
  createLoginWindow();

  // On OS X it's common to re-create a window in the app when the dock icon is clicked
  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
      createLoginWindow();
    }
  });
});


// Listen for login success event
ipcMain.on('login-success', (event, data) => {
  if (loginWindow) {
    loginWindow.close();
  }
  createMainWindow();

  // Send data to the main window if needed
  mainWindow.webContents.on('did-finish-load', () => {
    mainWindow.webContents.send('login-data', data);
  });
});
