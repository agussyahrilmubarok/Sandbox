import { Link as RouterLink, useNavigate } from "react-router-dom";
import {
  Flex,
  Box,
  Spacer,
  Link,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  Image,
} from "@chakra-ui/react";
import toast from "react-hot-toast";
import { useUser } from "../context/UserContext";
import { SERVER_BASE_URL, API_BASE_URL_V1 } from "../util.js";

export default function NavBar() {
  const { user, updateUser } = useUser();
  const navigate = useNavigate();
  const avatarUrl = user?.avatar ? `${SERVER_BASE_URL}${user.avatar}` : '/images/users/default-avatar.png';

  console.log(user);

  const handleSignOut = async () => {
    try {
      const res = await fetch(`${API_BASE_URL_V1}/auth/signout`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Failed to sign out");
      const data = await res.json();
      toast.success(data.message);
      updateUser(null);
      navigate("/");
    } catch (error) {
      toast.error(error.message || "Failed to sign out");
    }
  };

  return (
    <Box as="nav" bg="red.50">
      <Flex
        minWidth="max-content"
        alignItems="center"
        p="12px"
        maxW="7xl"
        m="0 auto"
      >
        <Box p="2">
          <Link as={RouterLink} fontSize="lg" fontWeight="bold" to="/">
            Taskly
          </Link>
        </Box>
        <Spacer />
        <Box>
          {user ? (
            <Menu>
              <MenuButton>
                <Image
                  boxSize="40px"
                  borderRadius="full"
                  src={avatarUrl}
                  alt={user.username}
                />
              </MenuButton>
              <MenuList>
                <MenuItem as={RouterLink} to="/profile">
                  Profile
                </MenuItem>
                <MenuItem onClick={handleSignOut}>Sign Out</MenuItem>
              </MenuList>
            </Menu>
          ) : (
            <Link as={RouterLink} to="/signin">
              Sign In
            </Link>
          )}
        </Box>
      </Flex>
    </Box>
  );
}
