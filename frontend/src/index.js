import React from 'react';
import {createRoot} from 'react-dom/client';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import 'semantic-ui-css/semantic.min.css';

const rootElement = document.getElementById('root');
const root = createRoot(rootElement);
root.render(<App />);
registerServiceWorker();