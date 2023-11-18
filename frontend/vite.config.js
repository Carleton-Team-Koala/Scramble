import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
/**
 * Vite configuration file.
 * @module vite.config
 */

/**
 * Vite configuration object.
 * @typedef {Object} ViteConfig
 * @property {Array} plugins - List of plugins to use.
 * @property {Object} server - Server configuration.
 * @property {string} server.host - Server host.
 * @property {number} server.port - Server port.
 */

export default defineConfig({
  /**
   * List of plugins to use.
   * @type {Array}
   */
  plugins: [react()],

  /**
   * Server configuration.
   * @type {Object}
   */
  server: {
    /**
     * Server host.
     * @type {string}
     */
    host: "0.0.0.0",

    /**
     * Server port.
     * @type {number}
     */
    port: 3000
  }
})
