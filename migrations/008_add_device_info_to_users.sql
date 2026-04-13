-- Migration: Add device fingerprinting columns to users table
-- These columns are populated automatically on every login from the HTTP request headers.

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS last_ip       VARCHAR(45),       -- IPv4 (max 15) or IPv6 (max 45)
    ADD COLUMN IF NOT EXISTS last_os       VARCHAR(100),      -- e.g. "Windows 10/11", "macOS"
    ADD COLUMN IF NOT EXISTS last_browser  VARCHAR(100),      -- e.g. "Google Chrome", "Mozilla Firefox"
    ADD COLUMN IF NOT EXISTS architecture  VARCHAR(20);       -- "64-bit", "32-bit", or "Unknown"

COMMENT ON COLUMN users.last_ip      IS 'IP address of the client on their last successful login';
COMMENT ON COLUMN users.last_os      IS 'Operating system detected from User-Agent on last login';
COMMENT ON COLUMN users.last_browser IS 'Browser detected from User-Agent on last login';
COMMENT ON COLUMN users.architecture IS 'CPU architecture detected from User-Agent on last login';
