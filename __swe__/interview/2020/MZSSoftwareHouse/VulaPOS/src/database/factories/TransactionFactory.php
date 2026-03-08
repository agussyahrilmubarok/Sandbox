<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Transaction>
 */
class TransactionFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'invoice' => fake()->unique()->numerify('INV-#####'),
            'cash' => fake()->numberBetween(1000, 100000),
            'change' => fake()->numberBetween(0, 5000),
            'discount' => fake()->numberBetween(0, 10000),
            'grand_total' => fake()->numberBetween(1000, 100000),
            'cashier_id' => \App\Models\User::factory(),
            'customer_id' => \App\Models\Customer::factory(),
        ];
    }
}
