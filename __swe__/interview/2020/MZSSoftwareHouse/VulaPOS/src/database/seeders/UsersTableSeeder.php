<?php

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use Spatie\Permission\Models\Role;
use Spatie\Permission\Models\Permission;

class UsersTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $this->makeAdmin();
    }

    private function makeAdmin(): void
    {
        $user = User::create([
            'name'      => 'Administrator',
            'email'     => 'admin@vulapos.com',
            'password'  => bcrypt('adminpanel123'),
        ]);

        $permissions = Permission::all();
        $role = Role::find(1);
        $role->syncPermissions($permissions);
        $user->assignRole($role);
    }
}
