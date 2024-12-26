/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.12.4
 * source: game.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as pb_1 from "google-protobuf";
export namespace game {
    export class Vector2 extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            x?: number;
            y?: number;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("x" in data && data.x != undefined) {
                    this.x = data.x;
                }
                if ("y" in data && data.y != undefined) {
                    this.y = data.y;
                }
            }
        }
        get x() {
            return pb_1.Message.getFieldWithDefault(this, 1, 0) as number;
        }
        set x(value: number) {
            pb_1.Message.setField(this, 1, value);
        }
        get y() {
            return pb_1.Message.getFieldWithDefault(this, 2, 0) as number;
        }
        set y(value: number) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            x?: number;
            y?: number;
        }): Vector2 {
            const message = new Vector2({});
            if (data.x != null) {
                message.x = data.x;
            }
            if (data.y != null) {
                message.y = data.y;
            }
            return message;
        }
        toObject() {
            const data: {
                x?: number;
                y?: number;
            } = {};
            if (this.x != null) {
                data.x = this.x;
            }
            if (this.y != null) {
                data.y = this.y;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.x != 0)
                writer.writeDouble(1, this.x);
            if (this.y != 0)
                writer.writeDouble(2, this.y);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Vector2 {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Vector2();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.x = reader.readDouble();
                        break;
                    case 2:
                        message.y = reader.readDouble();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Vector2 {
            return Vector2.deserialize(bytes);
        }
    }
    export class Controls extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            speed?: number;
            jump_force?: number;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("speed" in data && data.speed != undefined) {
                    this.speed = data.speed;
                }
                if ("jump_force" in data && data.jump_force != undefined) {
                    this.jump_force = data.jump_force;
                }
            }
        }
        get speed() {
            return pb_1.Message.getFieldWithDefault(this, 1, 0) as number;
        }
        set speed(value: number) {
            pb_1.Message.setField(this, 1, value);
        }
        get jump_force() {
            return pb_1.Message.getFieldWithDefault(this, 2, 0) as number;
        }
        set jump_force(value: number) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            speed?: number;
            jump_force?: number;
        }): Controls {
            const message = new Controls({});
            if (data.speed != null) {
                message.speed = data.speed;
            }
            if (data.jump_force != null) {
                message.jump_force = data.jump_force;
            }
            return message;
        }
        toObject() {
            const data: {
                speed?: number;
                jump_force?: number;
            } = {};
            if (this.speed != null) {
                data.speed = this.speed;
            }
            if (this.jump_force != null) {
                data.jump_force = this.jump_force;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.speed != 0)
                writer.writeDouble(1, this.speed);
            if (this.jump_force != 0)
                writer.writeDouble(2, this.jump_force);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Controls {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Controls();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.speed = reader.readDouble();
                        break;
                    case 2:
                        message.jump_force = reader.readDouble();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Controls {
            return Controls.deserialize(bytes);
        }
    }
    export class InputState extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            move_left?: boolean;
            move_right?: boolean;
            jump?: boolean;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("move_left" in data && data.move_left != undefined) {
                    this.move_left = data.move_left;
                }
                if ("move_right" in data && data.move_right != undefined) {
                    this.move_right = data.move_right;
                }
                if ("jump" in data && data.jump != undefined) {
                    this.jump = data.jump;
                }
            }
        }
        get move_left() {
            return pb_1.Message.getFieldWithDefault(this, 1, false) as boolean;
        }
        set move_left(value: boolean) {
            pb_1.Message.setField(this, 1, value);
        }
        get move_right() {
            return pb_1.Message.getFieldWithDefault(this, 2, false) as boolean;
        }
        set move_right(value: boolean) {
            pb_1.Message.setField(this, 2, value);
        }
        get jump() {
            return pb_1.Message.getFieldWithDefault(this, 3, false) as boolean;
        }
        set jump(value: boolean) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            move_left?: boolean;
            move_right?: boolean;
            jump?: boolean;
        }): InputState {
            const message = new InputState({});
            if (data.move_left != null) {
                message.move_left = data.move_left;
            }
            if (data.move_right != null) {
                message.move_right = data.move_right;
            }
            if (data.jump != null) {
                message.jump = data.jump;
            }
            return message;
        }
        toObject() {
            const data: {
                move_left?: boolean;
                move_right?: boolean;
                jump?: boolean;
            } = {};
            if (this.move_left != null) {
                data.move_left = this.move_left;
            }
            if (this.move_right != null) {
                data.move_right = this.move_right;
            }
            if (this.jump != null) {
                data.jump = this.jump;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.move_left != false)
                writer.writeBool(1, this.move_left);
            if (this.move_right != false)
                writer.writeBool(2, this.move_right);
            if (this.jump != false)
                writer.writeBool(3, this.jump);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): InputState {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new InputState();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.move_left = reader.readBool();
                        break;
                    case 2:
                        message.move_right = reader.readBool();
                        break;
                    case 3:
                        message.jump = reader.readBool();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): InputState {
            return InputState.deserialize(bytes);
        }
    }
    export class GameObject extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            id?: string;
            position?: Vector2;
            velocity?: Vector2;
            acceleration?: Vector2;
            mass?: number;
            elasticity?: number;
            friction?: number;
            controls?: Controls;
            Grounded?: boolean;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("id" in data && data.id != undefined) {
                    this.id = data.id;
                }
                if ("position" in data && data.position != undefined) {
                    this.position = data.position;
                }
                if ("velocity" in data && data.velocity != undefined) {
                    this.velocity = data.velocity;
                }
                if ("acceleration" in data && data.acceleration != undefined) {
                    this.acceleration = data.acceleration;
                }
                if ("mass" in data && data.mass != undefined) {
                    this.mass = data.mass;
                }
                if ("elasticity" in data && data.elasticity != undefined) {
                    this.elasticity = data.elasticity;
                }
                if ("friction" in data && data.friction != undefined) {
                    this.friction = data.friction;
                }
                if ("controls" in data && data.controls != undefined) {
                    this.controls = data.controls;
                }
                if ("Grounded" in data && data.Grounded != undefined) {
                    this.Grounded = data.Grounded;
                }
            }
        }
        get id() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set id(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get position() {
            return pb_1.Message.getWrapperField(this, Vector2, 2) as Vector2;
        }
        set position(value: Vector2) {
            pb_1.Message.setWrapperField(this, 2, value);
        }
        get has_position() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get velocity() {
            return pb_1.Message.getWrapperField(this, Vector2, 3) as Vector2;
        }
        set velocity(value: Vector2) {
            pb_1.Message.setWrapperField(this, 3, value);
        }
        get has_velocity() {
            return pb_1.Message.getField(this, 3) != null;
        }
        get acceleration() {
            return pb_1.Message.getWrapperField(this, Vector2, 4) as Vector2;
        }
        set acceleration(value: Vector2) {
            pb_1.Message.setWrapperField(this, 4, value);
        }
        get has_acceleration() {
            return pb_1.Message.getField(this, 4) != null;
        }
        get mass() {
            return pb_1.Message.getFieldWithDefault(this, 5, 0) as number;
        }
        set mass(value: number) {
            pb_1.Message.setField(this, 5, value);
        }
        get elasticity() {
            return pb_1.Message.getFieldWithDefault(this, 6, 0) as number;
        }
        set elasticity(value: number) {
            pb_1.Message.setField(this, 6, value);
        }
        get friction() {
            return pb_1.Message.getFieldWithDefault(this, 7, 0) as number;
        }
        set friction(value: number) {
            pb_1.Message.setField(this, 7, value);
        }
        get controls() {
            return pb_1.Message.getWrapperField(this, Controls, 8) as Controls;
        }
        set controls(value: Controls) {
            pb_1.Message.setWrapperField(this, 8, value);
        }
        get has_controls() {
            return pb_1.Message.getField(this, 8) != null;
        }
        get Grounded() {
            return pb_1.Message.getFieldWithDefault(this, 9, false) as boolean;
        }
        set Grounded(value: boolean) {
            pb_1.Message.setField(this, 9, value);
        }
        static fromObject(data: {
            id?: string;
            position?: ReturnType<typeof Vector2.prototype.toObject>;
            velocity?: ReturnType<typeof Vector2.prototype.toObject>;
            acceleration?: ReturnType<typeof Vector2.prototype.toObject>;
            mass?: number;
            elasticity?: number;
            friction?: number;
            controls?: ReturnType<typeof Controls.prototype.toObject>;
            Grounded?: boolean;
        }): GameObject {
            const message = new GameObject({});
            if (data.id != null) {
                message.id = data.id;
            }
            if (data.position != null) {
                message.position = Vector2.fromObject(data.position);
            }
            if (data.velocity != null) {
                message.velocity = Vector2.fromObject(data.velocity);
            }
            if (data.acceleration != null) {
                message.acceleration = Vector2.fromObject(data.acceleration);
            }
            if (data.mass != null) {
                message.mass = data.mass;
            }
            if (data.elasticity != null) {
                message.elasticity = data.elasticity;
            }
            if (data.friction != null) {
                message.friction = data.friction;
            }
            if (data.controls != null) {
                message.controls = Controls.fromObject(data.controls);
            }
            if (data.Grounded != null) {
                message.Grounded = data.Grounded;
            }
            return message;
        }
        toObject() {
            const data: {
                id?: string;
                position?: ReturnType<typeof Vector2.prototype.toObject>;
                velocity?: ReturnType<typeof Vector2.prototype.toObject>;
                acceleration?: ReturnType<typeof Vector2.prototype.toObject>;
                mass?: number;
                elasticity?: number;
                friction?: number;
                controls?: ReturnType<typeof Controls.prototype.toObject>;
                Grounded?: boolean;
            } = {};
            if (this.id != null) {
                data.id = this.id;
            }
            if (this.position != null) {
                data.position = this.position.toObject();
            }
            if (this.velocity != null) {
                data.velocity = this.velocity.toObject();
            }
            if (this.acceleration != null) {
                data.acceleration = this.acceleration.toObject();
            }
            if (this.mass != null) {
                data.mass = this.mass;
            }
            if (this.elasticity != null) {
                data.elasticity = this.elasticity;
            }
            if (this.friction != null) {
                data.friction = this.friction;
            }
            if (this.controls != null) {
                data.controls = this.controls.toObject();
            }
            if (this.Grounded != null) {
                data.Grounded = this.Grounded;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.id.length)
                writer.writeString(1, this.id);
            if (this.has_position)
                writer.writeMessage(2, this.position, () => this.position.serialize(writer));
            if (this.has_velocity)
                writer.writeMessage(3, this.velocity, () => this.velocity.serialize(writer));
            if (this.has_acceleration)
                writer.writeMessage(4, this.acceleration, () => this.acceleration.serialize(writer));
            if (this.mass != 0)
                writer.writeDouble(5, this.mass);
            if (this.elasticity != 0)
                writer.writeDouble(6, this.elasticity);
            if (this.friction != 0)
                writer.writeDouble(7, this.friction);
            if (this.has_controls)
                writer.writeMessage(8, this.controls, () => this.controls.serialize(writer));
            if (this.Grounded != false)
                writer.writeBool(9, this.Grounded);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GameObject {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GameObject();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.id = reader.readString();
                        break;
                    case 2:
                        reader.readMessage(message.position, () => message.position = Vector2.deserialize(reader));
                        break;
                    case 3:
                        reader.readMessage(message.velocity, () => message.velocity = Vector2.deserialize(reader));
                        break;
                    case 4:
                        reader.readMessage(message.acceleration, () => message.acceleration = Vector2.deserialize(reader));
                        break;
                    case 5:
                        message.mass = reader.readDouble();
                        break;
                    case 6:
                        message.elasticity = reader.readDouble();
                        break;
                    case 7:
                        message.friction = reader.readDouble();
                        break;
                    case 8:
                        reader.readMessage(message.controls, () => message.controls = Controls.deserialize(reader));
                        break;
                    case 9:
                        message.Grounded = reader.readBool();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GameObject {
            return GameObject.deserialize(bytes);
        }
    }
    export class PhysicsEngineState extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            objects?: GameObject[];
            gravity?: Vector2;
            input?: InputState;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("objects" in data && data.objects != undefined) {
                    this.objects = data.objects;
                }
                if ("gravity" in data && data.gravity != undefined) {
                    this.gravity = data.gravity;
                }
                if ("input" in data && data.input != undefined) {
                    this.input = data.input;
                }
            }
        }
        get objects() {
            return pb_1.Message.getRepeatedWrapperField(this, GameObject, 1) as GameObject[];
        }
        set objects(value: GameObject[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        get gravity() {
            return pb_1.Message.getWrapperField(this, Vector2, 2) as Vector2;
        }
        set gravity(value: Vector2) {
            pb_1.Message.setWrapperField(this, 2, value);
        }
        get has_gravity() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get input() {
            return pb_1.Message.getWrapperField(this, InputState, 3) as InputState;
        }
        set input(value: InputState) {
            pb_1.Message.setWrapperField(this, 3, value);
        }
        get has_input() {
            return pb_1.Message.getField(this, 3) != null;
        }
        static fromObject(data: {
            objects?: ReturnType<typeof GameObject.prototype.toObject>[];
            gravity?: ReturnType<typeof Vector2.prototype.toObject>;
            input?: ReturnType<typeof InputState.prototype.toObject>;
        }): PhysicsEngineState {
            const message = new PhysicsEngineState({});
            if (data.objects != null) {
                message.objects = data.objects.map(item => GameObject.fromObject(item));
            }
            if (data.gravity != null) {
                message.gravity = Vector2.fromObject(data.gravity);
            }
            if (data.input != null) {
                message.input = InputState.fromObject(data.input);
            }
            return message;
        }
        toObject() {
            const data: {
                objects?: ReturnType<typeof GameObject.prototype.toObject>[];
                gravity?: ReturnType<typeof Vector2.prototype.toObject>;
                input?: ReturnType<typeof InputState.prototype.toObject>;
            } = {};
            if (this.objects != null) {
                data.objects = this.objects.map((item: GameObject) => item.toObject());
            }
            if (this.gravity != null) {
                data.gravity = this.gravity.toObject();
            }
            if (this.input != null) {
                data.input = this.input.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.objects.length)
                writer.writeRepeatedMessage(1, this.objects, (item: GameObject) => item.serialize(writer));
            if (this.has_gravity)
                writer.writeMessage(2, this.gravity, () => this.gravity.serialize(writer));
            if (this.has_input)
                writer.writeMessage(3, this.input, () => this.input.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): PhysicsEngineState {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new PhysicsEngineState();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.objects, () => pb_1.Message.addToRepeatedWrapperField(message, 1, GameObject.deserialize(reader), GameObject));
                        break;
                    case 2:
                        reader.readMessage(message.gravity, () => message.gravity = Vector2.deserialize(reader));
                        break;
                    case 3:
                        reader.readMessage(message.input, () => message.input = InputState.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): PhysicsEngineState {
            return PhysicsEngineState.deserialize(bytes);
        }
    }
}
