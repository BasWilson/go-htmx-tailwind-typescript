import 'htmx.org';
import { enableHotReload } from './hot-reload';

if (window.location.hostname === 'localhost')
    enableHotReload();