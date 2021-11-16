<?php

return [

    /*
    |--------------------------------------------------------------------------
    | Cross-Origin Resource Sharing (CORS) Configuration
    |--------------------------------------------------------------------------
    |
    | Here you may configure your settings for cross-origin resource sharing
    | or "CORS". This determines what cross-origin operations may execute
    | in web browsers. You are free to adjust these settings as needed.
    |
    | To learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
    |
    */

    'paths' => ['*', 'api/*','login'],
    // 'paths' => [
    //     'api/*',
    //     'login',
    //     // 'logout',
    //     // 'register',
    //     // 'user/password',
    //     // 'forgot-password',
    //     // 'reset-password',
    //     'sanctum/csrf-cookie',
    //     // 'user/profile-information',
    //     // 'email/verification-notification',
    //   ],

    'allowed_methods' => ['*'],

    'allowed_origins' => ['http://localhost:3000', env('SPA_URL')],

    'allowed_origins_patterns' => [],

    'allowed_headers' => ['*'],
      // 'Access-Control-Allow-Credentials', 'Access-Control-Allow-Origin ','x-xsrf-token'],

    'exposed_headers' => [],

    //   /*
    //      * Preflight request will respond with value for the max age header.
    //      */
    //     'max_age' => 60 * 60 * 24,
    //     // 'max_age' => 0,

    'supports_credentials' => true,

];
