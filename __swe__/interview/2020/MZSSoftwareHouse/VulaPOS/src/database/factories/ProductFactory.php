<?php

namespace Database\Factories;

use App\Models\Category;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Product>
 */
class ProductFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'title' => fake()->word(),
            'description' => fake()->paragraph(),
            'stock' => fake()->numberBetween(1, 1000),
            'buy_price' => fake()->randomFloat(2, 1, 1000),
            'sell_price' => fake()->randomFloat(2, 1, 1000),
            'image' => fake()->imageUrl(640, 480, 'products', true),
            'barcode' => fake()->numerify('##########'),
            'category_id' => Category::factory(),
        ];
    }
}
