<?php

namespace App\Http\Middleware;

use Illuminate\Http\Request;
use Inertia\Middleware;

class HandleInertiaRequests extends Middleware
{
    /**
     * The root template that is loaded on the first page visit.
     *
     * @var string
     */
    protected $rootView = 'app';

    /**
     * Determine the current asset version.
     */
    public function version(Request $request): string|null
    {
        return parent::version($request);
    }

    /**
     * Define the props that are shared by default.
     *
     * @return array<string, mixed>
     */
    public function share(Request $request): array
    {
        $user = $request->user();

        return [
            ...parent::share($request),
            'auth' => [
                'user' => $user,
                'permissions' => $user ? $user->getPermissionArray() : [],
            ],
            'route' => function () use ($request) {
                return [
                    'params' => $request->route()->parameters(),
                    'query' => $request->all(),
                ];
            },
            'session' => [
                'status'    => fn () => $request->session()->get('status'),
                'success'   => fn () => $request->session()->get('success'),
                'error'     => fn () => $request->session()->get('error')
            ],
        ];
    }
}
