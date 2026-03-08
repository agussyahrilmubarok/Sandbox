<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Spatie\Permission\Models\Permission;
use Inertia\Inertia;
use Inertia\Response;

class PermissionController extends Controller
{
    /**
     * Handle the incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function __invoke(Request $request): Response
    {
        $q = $request->input('q');

        $permissions = Permission::when($q, function ($query, $q) {
            $query->where('name', 'like', '%' . $q . '%');
        })->latest()->paginate(5);

        return Inertia::render('Permissions/Index', [
            'permissions' => $permissions
        ]);
    }
}
