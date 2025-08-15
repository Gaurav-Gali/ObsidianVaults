![Source](https://youtu.be/t8ane4BDyC8?t=164)

==**Note** : Use yarn for better compateblity with electron==

1. Creating an electron app
```bash
yarn create @quick-start/electron
```

```bash
// To install the dependencies
yarn install
// To run the elctron app
yarn dev
```

2. Creating some important folders
	1. src/main/lib
		1. This would contain the api's for interacting with OS.
	2. src/shared
		1. This would store the shared resource between main, preloaded and renderer
	3. renderer/src/components
	4. renderer/src/utils
	5. renderer/src/assets
	6. renderer/src/store
	7. renderer/src/hooks

3. Working and structuring of an electron app
![[electron_working.png]]

#### Modify the following files
1. src/main/index.ts
```ts
import { app, shell, BrowserWindow, ipcMain } from 'electron';
import { join } from 'path';
import { electronApp, optimizer, is } from '@electron-toolkit/utils';
import icon from '../../resources/icon.png?asset';

function createWindow(): void {
  // Create the browser window.
  const mainWindow = new BrowserWindow({
    width: 900,
    height: 670,
    show: false,
    autoHideMenuBar: true,
    ...(process.platform === 'linux' ? { icon } : {}),

	// CUSTOMIZATIONS
	center: true,
	title: "Easy Notes",
	vibrancy: "under-window",
	visualEffectState: "active",
	titleBarStyle: "hidden",
	trafficLightPosition : {x:15,y:10},
	// END CUSTOMIZATIONS

    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: true,
      contextIsolation: true,
    },
  });

  mainWindow.on('ready-to-show', () => {
    mainWindow.show();
  });

  mainWindow.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url);
    return { action: 'deny' };
  });

  // HMR for renderer based on electron-vite CLI.
  // Load the remote URL for development or the local HTML file for production.
  if (is.dev && process.env['ELECTRON_RENDERER_URL']) {
    mainWindow.loadURL(process.env['ELECTRON_RENDERER_URL']);
  } else {
    mainWindow.loadFile(join(__dirname, '../renderer/index.html'));
  }
}

// This method will be called when Electron has finished initialization
// and is ready to create browser windows.
app.whenReady().then(() => {
  // Set app user model ID for Windows
  electronApp.setAppUserModelId('com.electron');

  // Default open or close DevTools by F12 in development
  // and ignore CommandOrControl + R in production.
  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window);
  });

  // IPC test
  ipcMain.on('ping', () => console.log('pong'));

  createWindow();

  app.on('activate', () => {
    // On macOS, re-create a window when the dock icon is clicked and
    // there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

// Quit when all windows are closed, except on macOS.
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});
```

2. src/preload/index.ts
```ts
import { contextBridge } from 'electron';

if (!process.contextIsolated) {
  throw new Error('Context must be isolated to use this API');
}

try {
  contextBridge.exposeInMainWorld('context', {
    // TODO: Add your own methods here
  });
} catch (error) {
  console.error('Failed to expose API to renderer:', error);
}
```

3. src/preload/index.d.ts
```ts
import { ElectronAPI } from '@electron-toolkit/preload';

declare global {
  interface Window {
    context: ElectronAPI & {
      // TODO: Add custom methods exposed via contextBridge
    };
  }
}
```

4. tsconfig.web.json
```json
{
  "extends": "@electron-toolkit/tsconfig/tsconfig.web.json",
  "include": [
    "src/renderer/src/env.d.ts",
    "src/renderer/src/**/*",
    "src/renderer/src/**/*.tsx",
    "src/preload/*.d.ts",
    "src/shared/**/*"
  ],
  "compilerOptions": {
    "composite": true,
    "jsx": "react-jsx",
    "noUnusedLocals": false,
    "baseUrl": ".",
    "paths": {
      "@renderer/*": ["src/renderer/src/*"],
      "@shared/*": ["src/shared/*"],
      "@/*": ["src/renderer/src/*"]
    }
  }
}
```

5. tsconfig.node.json
```json
{
  "extends": "@electron-toolkit/tsconfig/tsconfig.node.json",
  "include": [
    "electron.vite.config.*",
    "src/main/**/*",
    "src/preload/**/*",
    "src/shared/**/*"
  ],
  "compilerOptions": {
    "composite": true,
    "types": ["electron-vite/node"],
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/main/*"],
      "@shared/*": ["src/shared/*"]
    }
  }
}
```

6. electron.vite.config.ts
```ts
import { resolve } from 'path';
import { defineConfig, externalizeDepsPlugin } from 'electron-vite';
import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  main: {
    plugins: [externalizeDepsPlugin()],
    resolve: {
      alias: {
        '@/lib': resolve('src/main/lib'),
        '@shared': resolve('src/shared')
      }
    }
  },
  preload: {
    plugins: [externalizeDepsPlugin()]
  },
  renderer: {
    assetsInclude: 'src/renderer/assets/**',
    resolve: {
      alias: {
        '@renderer': resolve('src/renderer/src'),
        '@shared': resolve('src/shared'),
        '@/hooks': resolve('src/renderer/src/hooks'),
        '@/assets': resolve('src/renderer/src/assets'),
        '@/store': resolve('src/renderer/src/store'),
        '@/components': resolve('src/renderer/src/components'), // Fixed typo
        '@/mocks': resolve('src/renderer/src/mocks')
      }
    },
    plugins: [react(), tailwindcss()]
  }
});
```

#### Setting up TailwindCSS and ShadCN
#### Tailwind CSS
[**Follow the official docs**](https://tailwindcss.com/docs/installation/using-vite)

Radix UI
[**Follow the official docs**](https://www.radix-ui.com/themes/docs/overview/getting-started)

