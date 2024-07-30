const { contextBridge, ipcRenderer } = require('electron');

contextBridge.exposeInMainWorld('electron', {
  sendLoginData: (username, apikey) => ipcRenderer.send('login-success', { username , apikey }),
  onLoginData: (callback) => ipcRenderer.on('login-data', (event, data) => callback(data)),
  getCurrentDateTime: () => new Date().toLocaleString(),
});
