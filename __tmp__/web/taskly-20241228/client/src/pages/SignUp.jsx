import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import {
  Box,
  Heading,
  Text,
  Flex,
  Stack,
  FormControl,
  Input,
  FormErrorMessage,
  Button,
} from "@chakra-ui/react";
import toast from "react-hot-toast";

import { API_BASE_URL_V1 } from "../util";
import { useUser } from "../context/UserContext";

export default function SignUp() {
  const navigate = useNavigate();
  const { updateUser } = useUser();

  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm();

  const doSubmit = async (values) => {
    try {
      const res = await fetch(`${API_BASE_URL_V1}/auth/signup`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",   // use this for get cookie from server and save in browser
        body: JSON.stringify(values),
      });

      const data = await res.json();
      console.log(data);

      if (res.status == 200) {
        toast.success("Sign Up Successful. You are now logged in");
        updateUser(data); // save into context api
        navigate("/profile");
      } else {
        toast.error(data.message);
      }
    } catch (error) {
      console.log(error);
      toast.error("Something went wrong");
    }
  };

  return (
    <Box p="3" maxW="lg" mx="auto">
      <Heading
        as="h1"
        textAlign="center"
        fontSize="3xl"
        fontWeight="semibold"
        my="7"
      >
        Create an Account
      </Heading>
      {/* form details... */}
      <form onSubmit={handleSubmit(doSubmit)}>
        <Stack gap="4">
          <FormControl isInvalid={errors.username}>
            <Input
              id="username"
              type="text"
              placeholder="Your username"
              {...register("username", { required: "Username is required" })}
            />
            <FormErrorMessage>
              {errors.username && errors.username.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={errors.email}>
            <Input
              id="email"
              type="email"
              placeholder="Your email"
              {...register("email", { required: "Email is required" })}
            />
            <FormErrorMessage>
              {errors.email && errors.email.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={errors.password}>
            <Input
              id="password"
              type="password"
              placeholder="Your password"
              {...register("password", { required: "Password is required" })}
            />
            <FormErrorMessage>
              {errors.password && errors.password.message}
            </FormErrorMessage>
          </FormControl>
          <Button
            type="submit"
            isLoading={isSubmitting}
            colorScheme="teal"
            textTransform="uppercase"
          >
            Sign Up
          </Button>
        </Stack>
      </form>
      {/* end form details... */}
      <Flex gap="2" mt="5">
        <Text>Have an account?</Text>
        <Link to={"/signin"}>
          <Text as="span" color="blue.400">
            Sign in
          </Text>
        </Link>
      </Flex>
    </Box>
  );
}
