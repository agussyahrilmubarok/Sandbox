<?php

namespace Database\Factories;

use App\Models\Transaction;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Profit>
 */
class ProfitFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'total' => fake()->randomFloat(2, 10, 1000), // Random total with 2 decimal places between 10 and 1000
            'transaction_id' => Transaction::factory(), // Create a Transaction record for transaction_id
        ];
    }
}
