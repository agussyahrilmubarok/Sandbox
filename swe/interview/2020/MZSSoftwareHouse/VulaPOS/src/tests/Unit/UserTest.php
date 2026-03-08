<?php

namespace Tests\Unit;

use App\Models\User;
use Illuminate\Support\Facades\Hash;
use PHPUnit\Framework\Attributes\Test;
use Tests\TestCase;

class UserTest extends TestCase
{
    #[Test]
    public function test_user_factory_creates_valid_user()
    {
        $user = User::factory()->create();

        $this->assertNotNull($user->name);
        $this->assertNotNull($user->email);
        $this->assertNotNull($user->email_verified_at);
        $this->assertTrue(Hash::check('password', $user->password));
        $this->assertNotNull($user->remember_token);
        $this->assertEquals(10, strlen($user->remember_token));
    }

    #[Test]
    public function test_user_factory_unverified_state()
    {
        $user = User::factory()->unverified()->create();

        $this->assertNull($user->email_verified_at);
    }

    // #[Test]
    // public function test_user_factory_email_is_unique()
    // {
    //     $user1 = User::factory()->create(['email' => 'test@example.com']);
    //     $this->assertNotNull($user1);

    //     User::factory()->create(['email' => 'test@example.com']);
    //     $this->expectException(\Illuminate\Database\QueryException::class);
    // }

    #[Test]
    public function test_user_can_be_created_with_specific_password()
    {
        $password = 'specificpassword';
        $user = User::factory()->create(['password' => Hash::make($password)]);

        $this->assertTrue(Hash::check($password, $user->password));
    }

    #[Test]
    public function test_user_remember_token_is_string()
    {
        $user = User::factory()->create();
        $this->assertIsString($user->remember_token);
        $this->assertEquals(10, strlen($user->remember_token));
    }

    #[Test]
    public function test_user_email_can_be_verified()
    {
        $user = User::factory()->create(['email_verified_at' => null]);
        $this->assertNull($user->email_verified_at);

        $user->markEmailAsVerified();

        $this->assertNotNull($user->email_verified_at);
        $this->assertInstanceOf(\Illuminate\Support\Carbon::class, $user->email_verified_at);
    }

    #[Test]
    public function test_user_password_can_be_reset()
    {
        $user = User::factory()->create();
        $newPassword = 'newpassword';

        $user->password = Hash::make($newPassword);
        $user->save();

        $this->assertTrue(Hash::check($newPassword, $user->password));
    }

}
