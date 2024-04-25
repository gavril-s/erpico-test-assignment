-- Organization Table
CREATE TABLE organization (
  id SERIAL PRIMARY KEY
);

-- User Table
CREATE TABLE "user" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  last_name VARCHAR(128),
  phone VARCHAR(20),
  email VARCHAR(255),
  creator_id INT NOT NULL,
  trainer_id INT,
  number VARCHAR(45),
  pin VARCHAR(45),
  FOREIGN KEY (creator_id) REFERENCES "user"(id),
  FOREIGN KEY (trainer_id) REFERENCES "user"(id)
);

-- Activity Table
CREATE TABLE activity (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  title VARCHAR(255),
  external_id VARCHAR(255)
);

-- Activity Details Table
CREATE TABLE activity_details (
  id SERIAL PRIMARY KEY,
  activity_id INT NOT NULL,
  comment TEXT,
  duration INT,
  max_participants INT,
  min_participants INT DEFAULT 0,
  quota INT,
  waitlist_active BOOLEAN,
  waitlist_limit_minutes INT,
  FOREIGN KEY (activity_id) REFERENCES activity(id)
);

-- Activity Visibility Table
CREATE TABLE activity_visibility (
  id SERIAL PRIMARY KEY,
  activity_id INT NOT NULL,
  deleted BOOLEAN,
  link VARCHAR(425),
  hide_name BOOLEAN DEFAULT FALSE,
  hide_on_web BOOLEAN DEFAULT FALSE,
  show_in_chess BOOLEAN DEFAULT TRUE,
  FOREIGN KEY (activity_id) REFERENCES activity(id)
);

-- Activity Requirements Table
CREATE TABLE activity_requirements (
  id SERIAL PRIMARY KEY,
  activity_id INT NOT NULL,
  is_for_trainer_enabled BOOLEAN DEFAULT FALSE,
  require_pay BOOLEAN DEFAULT FALSE,
  require_pay_timeout INT,
  type VARCHAR(255) NOT NULL DEFAULT 'gym', -- Change ENUM to VARCHAR
  is_profile_required BOOLEAN DEFAULT FALSE,
  is_comment_required BOOLEAN DEFAULT FALSE,
  is_verification_required BOOLEAN,
  products_only_for_current_gym BOOLEAN,
  FOREIGN KEY (activity_id) REFERENCES activity(id)
);

-- Activity Additional Info Table
CREATE TABLE activity_additional_info (
  id SERIAL PRIMARY KEY,
  activity_id INT NOT NULL,
  comment TEXT,
  pack_cost INT NOT NULL DEFAULT 1,
  sequence INT,
  properties JSONB,
  FOREIGN KEY (activity_id) REFERENCES activity(id)
);

-- Schedule Table
CREATE TABLE schedule (
  id SERIAL PRIMARY KEY,
  org_id INT NOT NULL DEFAULT 1,
  activity_id INT,
  trainer_id INT,
  external_id VARCHAR(255),
  name TEXT,
  description TEXT,
  FOREIGN KEY (activity_id) REFERENCES activity(id),
  FOREIGN KEY (trainer_id) REFERENCES "user"(id)
);

-- Schedule Details Table
CREATE TABLE schedule_details (
  id SERIAL PRIMARY KEY,
  schedule_id INT NOT NULL,
  start_time TIME,
  end_time TIME,
  activity_date DATE,
  start_date DATE,
  activity_duration SMALLINT,
  cycle_day SMALLINT,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

-- Schedule Capacity Participation Table
CREATE TABLE schedule_capacity_participation (
  id SERIAL PRIMARY KEY,
  schedule_id INT NOT NULL,
  max_participants INT,
  min_participants INT DEFAULT 0,
  waitlist_active BOOLEAN,
  client_can_cancel BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

-- Schedule Type Visibility Table
CREATE TABLE schedule_type_visibility (
  id SERIAL PRIMARY KEY,
  schedule_id INT NOT NULL,
  type SMALLINT DEFAULT 1,
  deleted BOOLEAN DEFAULT FALSE,
  visible BOOLEAN DEFAULT TRUE,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

-- Schedule Date Activation Deactivation Table
CREATE TABLE schedule_date_activation_deactivation (
  id SERIAL PRIMARY KEY,
  schedule_id INT NOT NULL,
  date_activate DATE,
  date_deactivate DATE,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

-- Schedule Location Resources Table
CREATE TABLE schedule_location_resources (
  id SERIAL PRIMARY KEY,
  schedule_id INT NOT NULL,
  room_id INT,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

-- Schedule Additional Info Table
CREATE TABLE schedule_additional_info (
  id SERIAL PRIMARY KEY,
  schedule_id INT NOT NULL,
  video_link TEXT,
  sequence INT,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

-- Record Table
CREATE TABLE record (
  id SERIAL PRIMARY KEY,
  user_id INT,
  schedule_id INT,
  activity_id INT,
  activity_date DATE,
  start_time TIME,
  comment TEXT,
  state VARCHAR(255) DEFAULT 'draft', -- Change ENUM to VARCHAR
  parent_id INT,
  overlapped_record_id INT,
  FOREIGN KEY (user_id) REFERENCES "user"(id),
  FOREIGN KEY (schedule_id) REFERENCES schedule(id),
  FOREIGN KEY (activity_id) REFERENCES activity(id)
);

-- Record Equipment Table
CREATE TABLE record_equipment (
  record_id INT NOT NULL,
  equipment_id INT NOT NULL,
  deleted BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (record_id, equipment_id)
);
